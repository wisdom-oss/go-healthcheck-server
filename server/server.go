/*
 * Copyright (c) 2024, licensed under the EUPL-1.2-or-later
 */

// Package server provides a healthcheck server which allows running health
// checks directly in the main application. It is compatible with Unix-based
// operating systems and Windows operating systems
package server

// This file contains the code that can be shared between Unix-based and windows
// operating systems as the code written here is platform independent

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/wisdom-oss/go-healthcheck/common"
)

// HealthcheckServer represents a server responsible for handling health checks.
// Before using it, it needs to be initialized and a healthcheck function needs
// to be declared and set.
// Afterward you may start the healthcheck server with the Start function and
// wait for connections in a goroutine using the Run function
type HealthcheckServer struct {
	socketName          string
	listener            net.Listener
	ctx                 context.Context
	cancel              context.CancelCauseFunc
	healthcheckFunction func() error
}

// generateSocketName generates a random socket name and assigns it to the
// socketName field of the HealthcheckServer struct.
func (s *HealthcheckServer) generateSocketName() {
	s.socketName = randomString(16)
	f, err := os.Create(common.HealthcheckSocketNameFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(s.socketName)
	if err != nil {
		panic(err)
	}
}

// Init initializes the healthcheck server with a default context
func (s *HealthcheckServer) Init() {
	s.ctx, s.cancel = context.WithCancelCause(context.Background())
	s.generateSocketName()
}

// InitWithFunc initializes the healthcheck server with a default context and
// allows specifying the healthcheck function directly
func (s *HealthcheckServer) InitWithFunc(f func() error) {
	s.ctx, s.cancel = context.WithCancelCause(context.Background())
	s.healthcheckFunction = f
	s.generateSocketName()
}

// InitWithContext initializes the healthcheck server with a custom context
func (s *HealthcheckServer) InitWithContext(ctx context.Context) {
	s.ctx, s.cancel = context.WithCancelCause(ctx)
	s.generateSocketName()
}

// InitWithCancellableContext initializes the healthcheck server with a custom
// context and cancel function
func (s *HealthcheckServer) InitWithCancellableContext(ctx context.Context, cancelFunc context.CancelCauseFunc) {
	s.ctx = ctx
	s.cancel = cancelFunc
	s.generateSocketName()
}

// InitFull initializes the healthcheck server with the healthcheck function
// and cancellable context directly
func (s *HealthcheckServer) InitFull(f func() error, ctx context.Context, cancelFunc context.CancelCauseFunc) {
	s.ctx = ctx
	s.cancel = cancelFunc
	s.healthcheckFunction = f
	s.generateSocketName()
}

// Run begins listening for incoming connections on the HealthcheckServer's
// listener.
// It continuously accepts connections and launches a goroutine to handle each
// connection.
// The goroutine reads the input from the connection and checks if it matches
// the healthcheck command.
// If the command is matched, it calls the healthcheck function and writes the
// healthcheck response to the connection.
// If any errors occur during the process, the goroutine writes the error
// message to the connection and closes the connection.
// The function will exit when the context is canceled, which happens when the
// Stop method is called on the HealthcheckServer.
//
// ## Usage
//
// Since the Run function is blocking it is recommended to run in a goroutine
func (s *HealthcheckServer) Run() {
	defer s.listener.Close()
	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			conn, err := s.listener.Accept()
			if err != nil {
				s.cancel(err)
				continue
			}
			go func(connection net.Conn) {
				defer connection.Close()
				connection.SetReadDeadline(time.Now().Add(10 * time.Second))
				inputBuf := make([]byte, common.BufferSize)
				byteCount, err := connection.Read(inputBuf)
				if err != nil {
					s.cancel(err)
					return
				}
				message := strings.TrimSpace(string(inputBuf[:byteCount]))
				if message != common.HealthcheckIPCCommand {
					connection.Write([]byte(fmt.Sprintf("send '%s' to trigger healthcheck", common.HealthcheckIPCCommand)))
					return
				}
				err = s.healthcheckFunction()
				if err != nil {
					connection.Write([]byte(err.Error()))
					return
				}
				connection.Write([]byte(common.HealthcheckIPCResponse))
			}(conn)
		}
	}
}

// Stop cancels the context of the healthcheck server which results in the Run
// function to be stopped and the socket closing automatically
func (s *HealthcheckServer) Stop() {
	s.cancel(nil)
}
