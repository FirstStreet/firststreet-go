package backend

import (
	// "net/url"
	"log"
	"net/http"
	"sync"
	// "testing"
)

var (
	parseResponseOnce sync.Once
)

func parseResponseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.FormValue("token")
	log.Println(token)
	if token == "" {
		w.Write([]byte(`{"status": "OK", "message": "Of course I still love you"}`))
		return
	}
	response := []byte(`{"ok": true}`)
	w.Write(response)
}

func setParseResponseHandler() {
	http.HandleFunc("/parseResponse", parseResponseHandler)
}
