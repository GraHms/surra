# Surra Framework



Surra is a backend framework written in Go (Golang). It features a martini-like and Gin Framework, designed to easily develop TMF REST APIs


## Installation

To install Surra package, you need to install Go and set your Go workspace first.

1. You first need [Go](https://golang.org/) installed (**version 1.16+ is required**), then you can use the below Go command to install Surra.

```sh
$ go get -u github.com/grahms/surra
```

2. Import it in your code:

```go
import "github.com/grahms/surra"
```

3. (Optional) Import `net/http`. This is required for example if using constants such as `http.StatusOK`.

```go
import "net/http"
```

## Quick start

```sh
# assume the following codes in example.go file
$ cat app.go
```

```go
package main

import (
	"github.com/grahms/surra"
	"net/http"
)


func main() {
	r := surra.Principal()
	r.GET("/hello", func(ctx *surra.Context) {
		ctx.JSON(http.StatusOK, surra.KeyVal{
			"message": "pong",
		})
	})
	r.Run()

}
```

```
# run example.go and visit 0.0.0.0:8080/ping (for windows "localhost:8080/ping") on browser
$ go run example.go
```
