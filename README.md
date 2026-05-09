ginpprof
========

[![Go Reference](https://pkg.go.dev/badge/github.com/DeanThompson/ginpprof.svg)](https://pkg.go.dev/github.com/DeanThompson/ginpprof)

A wrapper for [gin](https://github.com/gin-gonic/gin) to mount `net/http/pprof` routes quickly.

## Why this project

`ginpprof` was created before `gin-contrib/pprof`, and many existing services still rely on it.
This repository is maintained to keep those users on a stable and lightweight API.

## Install

```sh
go get github.com/DeanThompson/ginpprof@latest
```

## Usage

```go
package main

import (
	"github.com/gin-gonic/gin"

	"github.com/DeanThompson/ginpprof"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Mount default pprof routes under /debug/pprof
	ginpprof.Wrap(router)

	// Or mount via group:
	// group := router.Group("/debug/pprof")
	// ginpprof.WrapGroup(group)

	router.Run(":8080")
}
```

### Registered routes

- `GET /debug/pprof/`
- `GET /debug/pprof/heap`
- `GET /debug/pprof/goroutine`
- `GET /debug/pprof/allocs`
- `GET /debug/pprof/block`
- `GET /debug/pprof/threadcreate`
- `GET /debug/pprof/cmdline`
- `GET /debug/pprof/profile`
- `GET /debug/pprof/symbol`
- `POST /debug/pprof/symbol`
- `GET /debug/pprof/trace`
- `GET /debug/pprof/mutex`

## Security notes

`pprof` endpoints expose runtime internals. For production:

- protect routes with auth middleware,
- restrict by network/IP,
- or expose only in internal environments.

## Development

```sh
go test ./...
```
