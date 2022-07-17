package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func health(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Print(checkHealth())
	response, _ := json.Marshal(checkHealth())
	fmt.Fprintf(w, string(response))
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/health", health)

	http.ListenAndServe(":8000", nil)
}