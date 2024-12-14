/*
 * Copyright (c) 2024, licensed under the EUPL-1.2-or-later
 */

package server

import (
	"fmt"
	"net"
	"os"

	"github.com/wisdom-oss/go-healthcheck/common"
)

// This file contains the unix-related implementation for the healthcheck
// server. It uses basic unix sockets and the packages provided in the std
// libraries

// Start starts the healthcheck server by listening on the configured socket.
// It initializes the `listener` field and returns an error if it fails to start
// the server.
func (s *HealthcheckServer) Start() (err error) {
	os.MkdirAll(common.SocketPrefix, os.ModeDir)
	if _, err = os.Create(common.SocketPrefix + s.socketName + ".sock"); err != nil {
		return err
	}
	s.listener, err = net.Listen("unix", common.SocketPrefix+s.socketName+".sock")
	if err != nil {
		return fmt.Errorf("unable to start healthcheck server: %w", err)
	}
	return nil
}
