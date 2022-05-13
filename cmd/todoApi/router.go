package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// InitialiseRouter initialisation des routes du web service
func InitialiseRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router = CORS(router)
	return router
}

// CORS ...
func CORS(router *mux.Router) *mux.Router {

	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type, Access-Token")
		w.WriteHeader(http.StatusOK)
	})
	return router
}
