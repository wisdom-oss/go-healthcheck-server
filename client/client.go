// Package client provides a client which automatically detects the healthcheck
// server and connects itself to the healthcheck server and triggers the
// healthcheck and awaits the results.
//
// ## Usage
//
//	import _ "github.com/wisdom-oss/go-healthcheck/client"
package client

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/wisdom-oss/go-healthcheck/common"
)

const healthcheckFlag = "-healthcheck"

func init() {
	// get the arguments the main executable has been called with but exclude
	// the executable name
	arguments := os.Args[1:]

	// now check if the arguments contain the healthcheck flag. if not exit the
	// init function and let the code continue as expected
	if !slices.Contains(arguments, healthcheckFlag) {
		return
	}

	// since the healthcheck flag has been set, try to find the socket name
	// in the environment
	socketNameFile, err := os.Open(common.HealthcheckSocketNameFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open socket name file: %v", err)
		os.Exit(1)
	}
	defer socketNameFile.Close()

	// now validate that the string is not empty
	socketNameBytes, err := io.ReadAll(socketNameFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read socket name: %v", err)
		os.Exit(1)
	}
	socketName := string(socketNameBytes)
	if strings.TrimSpace(socketName) == "" {
		fmt.Fprintf(os.Stderr, "empty socket name set in '%s'", common.HealthcheckSocketNameFile)
		os.Exit(1)
	}

	// now connect to the socket
	conn, err := connectSocket(socketName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to socket: %v", err)
		os.Exit(1)
	}
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	// now write the healthcheck command to the socket
	_, err = conn.Write([]byte(common.HealthcheckIPCCommand))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to write healthcheck command to socket: %v", err)
		os.Exit(1)
	}
	// now create an input buffer for checking the returned message
	inputBuf := make([]byte, common.BufferSize)
	byteCount, err := conn.Read(inputBuf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read healthcheck response from socket: %v", err)
		os.Exit(1)
	}
	// now read the response into a string
	response := strings.TrimSpace(string(inputBuf[:byteCount]))

	// now check if it is the expected response
	if response != common.HealthcheckIPCResponse {
		// since the server did not report successful healthcheck print the
		// response received and exit with 1
		fmt.Fprint(os.Stderr, response+"\n")
		os.Exit(1)
	}

	// since the healthcheck passed, exit with 0
	fmt.Fprint(os.Stdout, "healthcheck passed. service operational")
	os.Exit(0)
}
