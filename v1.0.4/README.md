<div align="center">
<img height="150px" src="https://raw.githubusercontent.com/wisdom-oss/brand/main/svg/standalone_color.svg">
<h1>Healthcheck Server/Client</h1>
<h3>go-healthcheck</h3>
<p>🫀Healthcheck server/client for microservices</p>
<img src="https://img.shields.io/github/go-mod/go-version/wisdom-oss/go-healthcheck?style=for-the-badge" alt="Go Lang Version"/>
<a href="https://pkg.go.dev/github.com/wisdom-oss/go-healthcheck">
<img alt="Go Reference" src="https://img.shields.io/badge/reference-reference?style=for-the-badge&logo=go&logoColor=white&labelColor=5c5c5c&color=287d9d">
</a>
</div>

## About
This package provides a healthcheck server and client which utilize UNIX sockets
on UNIX-like operating systems and named pipes on Windows operating systems.
The server opens up a socket or a named pipe and waits for a message indicating
that a healthcheck should be executed.
The socket or pipe is randomly named, and the name is added to the environment
variables as `HEALTHCHECK_SOCKET`.
This value is then picked up by the client which allows a seamless integration
into already existing software.

## Specifying a healthcheck function
> [!IMPORTANT]
> Since the healthcheck function is executed in a goroutine, you need to check
> your code if running the concurrently is possible. 
> If not, please use a `sync.Mutex` to protect your code and healthcheck 
> function against race conditions. 
> To check for race conditions, you may run your code with the following 
> command:
> ```bash
> go run -race .
> ```

Since there is no possibility to create a single healthcheck function which fits
all use cases, the server awaits a healthcheck function returning an error which
uses the following function signature:
```go
func healthcheckFunction() (err error) {
	// TODO: place your healthcheck functions here
}
```

