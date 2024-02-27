<div align="center">
<img height="150px" src="https://raw.githubusercontent.com/wisdom-oss/brand/main/svg/standalone_color.svg">
<h1>Healthcheck Server</h1>
<h3>go-healthcheck</h3>
<p>🫀Healthcheck server for microservices utilizing UNIX sockets or named pipes</p>
<img src="https://img.shields.io/github/go-mod/go-version/wisdom-oss/go-healthcheck?style=for-the-badge" alt="Go Lang Version"/>
<a href="https://pkg.go.dev/github.com/wisdom-oss/go-healthcheck">
<img alt="Go Reference" src="https://img.shields.io/badge/reference-reference?style=for-the-badge&logo=go&logoColor=white&labelColor=5c5c5c&color=287d9d">
</a>
</div>

> [!IMPORTANT]
> This healthcheck server is only compatible with Windows and Linux operating
> systems

> [!NOTE]
> The term _socket_ is used throughout the documentation as a synonym for the
> _named pipes_ provided by Windows to improve readability

## About
This package supplies the healthcheck server from the `github.com/wisdom-oss/go-healthcheck`
package.
On Linux operating systems, the server uses the `unix` network and sockets to
achieve inter-process communication with the client.
Since the `unix` network and sockets are not available under Windows, the server
uses [named pipes] from the Windows IPC API using the `go-winio` module.

[named pipes]: https://learn.microsoft.com/en-us/windows/win32/ipc/named-pipes

## Configuring
### Socket Name/Pipe Name
The socket name is generated automatically at during the initialization
process. It then is written into the `.hc-socket` file in the working directory
which allows the [client](../client/README.md) to automatically pick up the
socket and connect to it.

### Healthcheck Function
Since the server is not equipped with a default healthcheck function you need to
supply your own function.
This is done to allow a high degree of customization since not every 
microservice or other software product requires the same healthcheck.
The function you need to provide uses the following function signature:
```go
func healthcheck() error {
	// TODO: Implement healthcheck
}
```
The function should return an error if something is not working correctly to
indicate the failure of the healthcheck.
The error message is built by calling `Error()` on the returned object and sent
back using the socket to the client.