package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func hostInfo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Print(checkHealth())
	response, _ := json.Marshal(NewHostInfo())
	fmt.Fprintf(w, string(response))
}

func health(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Print(checkHealth())
	response, _ := json.Marshal(checkHealth())
	fmt.Fprintf(w, string(response))
}

func main() {
	http.HandleFunc("/hostinfo", hostInfo)
	http.HandleFunc("/health", health)

	http.ListenAndServe(":8000", nil)
}
