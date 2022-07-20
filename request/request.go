package request

import (
	"fmt"
	"html"
	"net/http"
)

type Request struct {
}

func New() *Request {
	return &Request{}
}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
func POST(path string, callback func()) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
}
