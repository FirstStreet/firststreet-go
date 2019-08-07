// Package testutil provides testing mocks and funcs
package testutil

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
)

var (
	ServerAddr string
	Once       sync.Once
)

// Default handler will return data that resolves
func DefaultHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ground control to major tom")
	})
}

func StartServer(handler http.HandlerFunc) {
	server := httptest.NewServer(handler)
	ServerAddr = "http://" + server.Listener.Addr().String()
}
