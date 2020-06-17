![](https://github.com/claudiobizzotto/go-http-echo-server/workflows/ci-cd/badge.svg)

# Golang HTTP Echo Server

The dumbest HTTP server you can find. Send a request and have it echoed back to you.

## Usage

```go
go run echoserver.go
```

Then point your browser to [http://localhost:8080](http://localhost:8080).

The port will be set to `8080` by default. If you want to run the server on another port, pass it as a parameter:

```go
go run echoserver.go 8090
```

## Docker

Follow the instructions below to _dockerize_ the server.

### Build image

```bash
docker build -t go-http-echo-server .
```

If you don't want to use the default port (`8080`), make sure you change that in the Dockerfile before building the image.

### Run container

```bash
docker run -d -p 8080:8080 go-http-echo-server
```
