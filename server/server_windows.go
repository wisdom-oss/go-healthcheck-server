package server

import (
	"fmt"

	"github.com/Microsoft/go-winio"

	"github.com/wisdom-oss/go-healthcheck/common"
)

// This file contains the windows-related implementation for the healthcheck
// server. It uses [Windows Named Pipes] instead of unix sockets as they
// are unavailable on Windows
//
// [Windows Named Pipes]: https://learn.microsoft.com/en-us/windows/win32/ipc/named-pipes

// Start starts the healthcheck server by listening on the configured socket.
// It initializes the `listener` field and returns an error if it fails to start
// the server.
func (s *HealthcheckServer) Start() (err error) {
	s.listener, err = winio.ListenPipe(common.SocketPrefix+s.socketName, nil)
	if err != nil {
		return fmt.Errorf("unable to start healthcheck server: %w", err)
	}
	return nil
}
