package views

import (
	"encoding/json"
	"fmt"
	"github.com/Abashinos/otus-msa-hw/server/middleware"
	"github.com/Abashinos/otus-msa-hw/server/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func listUsers(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", 501)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Unable to decode the request body into User model. %v", err)
		http.Error(w, err.Error(), 400)
		return
	}

	db, err := middleware.CreateConnection()
	if err != nil {
		log.Printf("Database unavailable. %v", err)
		http.Error(w, err.Error(), 500)
		return
	}

	result := db.Create(&user)
	if result.Error != nil {
		log.Printf("Unable to create User instance. %v", err)
		http.Error(w, result.Error.Error(), 500)
		return
	}

	response, _ := json.Marshal(&user)
	fmt.Println(w, string(response))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", 501)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", 501)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", 501)
}

type UserSubrouter struct {
}

func (*UserSubrouter) AddRoutes(r *mux.Router, prefix string) {
	sr := r.PathPrefix(prefix).Subrouter().StrictSlash(false)

	sr.HandleFunc("/", listUsers).Methods("GET", "OPTIONS")
	sr.HandleFunc("/", createUser).Methods("POST", "OPTIONS")
	sr.HandleFunc("/{id}", getUser).Methods("GET", "OPTIONS")
	sr.HandleFunc("/{id}", updateUser).Methods("PUT", "OPTIONS")
	sr.HandleFunc("/{id}", deleteUser).Methods("DELETE", "OPTIONS")
}
