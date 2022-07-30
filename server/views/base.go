package views

import "github.com/gorilla/mux"

type Subrouter interface {
	AddRoutes(router *mux.Router, prefix string)
}
