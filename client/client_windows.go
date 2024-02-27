package client

import (
	"net"

	"github.com/Microsoft/go-winio"
)

// socketPrefix contains the prefix for the path under which the socket or named
// pipe is created
const socketPrefix = `\\.\pipe\`

// connectSocket wraps the call to [winio.DialPipe] to allow an
// operating-system-based usage of named pipes and native unix sockets
func connectSocket(name string) (net.Conn, error) {
	return winio.DialPipe(socketPrefix+name, nil)
}
