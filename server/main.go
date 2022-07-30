package main

import (
	"encoding/json"
	"fmt"
	"github.com/Abashinos/otus-msa-hw/server/util"
	"github.com/Abashinos/otus-msa-hw/server/views"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func hostInfo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Print(checkHealth())
	response, _ := json.Marshal(NewHostInfo())
	fmt.Fprintf(w, string(response))
}

func health(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(checkHealth())
	fmt.Fprintf(w, string(response))
}

func debug(w http.ResponseWriter, req *http.Request) {
	// TODO: template
	data := struct {
		Env map[string]string `json:"env"`
	}{
		Env: util.DumpEnv(),
	}
	response, _ := json.Marshal(data)
	fmt.Fprintf(w, string(response))
}

func main() {
	r := mux.NewRouter().StrictSlash(false)

	r.HandleFunc("/hostinfo", hostInfo)
	r.HandleFunc("/health", health)
	r.HandleFunc("/debug", debug)

	usr := &views.UserSubrouter{}
	usr.AddRoutes(r, "/users")

	portStr := util.GetEnv("SERVER_PORT", "8000")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(fmt.Sprintf("Illegal port value: %s", portStr))
	}
	log.Printf("Starting server on %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}
