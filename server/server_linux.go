package server

import (
	"fmt"
	"net"

	"github.com/wisdom-oss/go-healthcheck/common"
)

// This file contains the unix-related implementation for the healthcheck
// server. It uses basic unix sockets and the packages provided in the std
// libraries

// Start starts the healthcheck server by listening on the configured socket.
// It initializes the `listener` field and returns an error if it fails to start
// the server.
func (s *HealthcheckServer) Start() (err error) {
	s.listener, err = net.Listen("unix", common.SocketPrefix+s.socketName+".sock")
	if err != nil {
		return fmt.Errorf("unable to start healthcheck server: %w", err)
	}
	return nil
}
