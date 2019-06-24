// Package testutil provides testing mocks and funcs
package testutil

import (
	"net/http/httptest"
	"sync"
)

var (
	ServerAddr string
	Once       sync.Once
)

func StartServer() {
	server := httptest.NewServer(nil)
	ServerAddr = "http://" + server.Listener.Addr().String()
}
