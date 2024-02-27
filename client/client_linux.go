package client

import (
	"net"

	"github.com/wisdom-oss/go-healthcheck/common"
)

// connectSocket wraps the call to [net.Dial] to allow an
// operating-system-based usage of named pipes and native unix sockets
func connectSocket(name string) (net.Conn, error) {
	return net.Dial("unix", common.SocketPrefix+name+".sock")
}
