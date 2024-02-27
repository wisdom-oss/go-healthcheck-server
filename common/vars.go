// Package common provides constants shared between the client and the server
// to allow central management for these variables.
package common

// BufferSize dictates the size for the buffers used for reading on the server
// and client side
const BufferSize = 4096

// HealthcheckSocketNameFile contains the name of the file used to store and
// look up the socket name
const HealthcheckSocketNameFile = ".hc-socket"

// HealthcheckIPCCommand contains the message required to trigger a healthcheck
// on the server
const HealthcheckIPCCommand = "ping"

// HealthcheckIPCResponse is returned if the healthcheck has been successful and
// the microservice/software product seams to operate as expected
const HealthcheckIPCResponse = "healthy"
