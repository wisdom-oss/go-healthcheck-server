<div align="center">
<img height="150px" src="https://raw.githubusercontent.com/wisdom-oss/brand/main/svg/standalone_color.svg">
<h1>Healthcheck Client</h1>
<h3>go-healthcheck</h3>
<p>🫀Healthcheck client for microservices</p>
<img src="https://img.shields.io/github/go-mod/go-version/wisdom-oss/go-healthcheck?style=for-the-badge" alt="Go Lang Version"/>
<a href="https://pkg.go.dev/github.com/wisdom-oss/go-healthcheck">
<img alt="Go Reference" src="https://img.shields.io/badge/reference-reference?style=for-the-badge&logo=go&logoColor=white&labelColor=5c5c5c&color=287d9d">
</a>
</div>

> [!IMPORTANT]
> This healthcheck client is only compatible with Windows and Linux operating
> systems

> [!NOTE]
> The term _socket_ is used throughout the documentation as a synonym for the
> _named pipes_ provided by Windows to improve readability

## About
This package supplies the healthcheck client from the `github.com/wisdom-oss/go-healthcheck`
package.
On Linux operating systems, the client uses the `unix` network and sockets to
achieve inter-process communication with the client.
Since the `unix` network and sockets are not available under Windows, the client
uses [named pipes] from the Windows IPC API using the `go-winio` module.

[named pipes]: https://learn.microsoft.com/en-us/windows/win32/ipc/named-pipes

## Configuring
### Socket Name
Since the [server] generates the socket name automatically, the client will try
to pick up the socket name from the `.hc-socket` file.

[server]: ../server/README.md

## Usage
To allow a seamless function of the healthcheck client, it only needs to be
imported for its side effects using the following import statement
```go
import _ github.com/wisdom-oss/go-healthcheck/client
```
Since the client uses a `init()` function to determine if the healthcheck should
be requested this is the only thing you need to do.
Afterward you may call your executable with the `-healthcheck` flag which will
trigger the healthcheck request. 

Since the targeted use case for the client is a docker container running 
automatic healthchecks, the response buffer size is limited to 4096 bytes as
this is the number of output bytes kept by Docker after a health check