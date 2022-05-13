package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter(r *mux.Router) {
	r.Methods("GET").Path("/health").HandlerFunc(health)
	r.Methods("GET").Path("/todos").HandlerFunc(GetTodos)
	r.Methods("PUT").Path("/todos").HandlerFunc(PutTodo)
	r.Methods("POST").Path("/todos/{id}/completed").HandlerFunc(TodoIsDone)
	r.Methods("DELETE").Path("/todos/{id}").HandlerFunc(DeleteTodo)
}

func health(w http.ResponseWriter, req *http.Request) {
	ReturnResponse(w, http.StatusOK, "Healthy")
}
