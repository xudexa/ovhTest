package app

import (
	"encoding/json"
	"errors"

	"net/http"

	"github.com/gorilla/mux"
)

func GetTodos(w http.ResponseWriter, req *http.Request) {
	var (
		response             WSResponse
		params               QueryParams
		todos, retTodos      []Todo
		totalCount, maxCount int
		status               int
		err                  error
	)

	// Recovering query parameters
	params = params.Parse(req)
	status = http.StatusOK

	todos = ListTodos(params)

	totalCount = len(todos)
	if totalCount == 0 {
		status = http.StatusNotFound
		err = errors.New(" Data not found. ")
		ReturnResponse(w, status, err)
	} else {
		meta := MetaResponse{ObjectName: "Todos"}

		low := params.Offset - 1
		if low == -1 {
			low = 0
		}

		// CountMax calculation available
		maxCount = params.Count
		if maxCount == 0 {
			maxCount = 100
		}

		high := maxCount + low
		if high > totalCount {
			high = totalCount
		}

		if low > high {
			status = http.StatusBadRequest
			err = errors.New(" Offset cannot be higher than count. ")
			ReturnResponse(w, status, err)
		}

		retTodos = todos[low:high]

		meta.TotalCount = len(todos)
		meta.Count = len(retTodos)
		meta.Offset = low + 1

		response.Meta = meta
		response.Data = retTodos

		ReturnResponse(w, status, response)
	}

}

func PutTodo(w http.ResponseWriter, req *http.Request) {
	var (
		params QueryParams

		todo Todo
	)

	params = params.Parse(req)
	json.Unmarshal(params.Body, &todo)
	AddTodo(todo)
	ReturnResponse(w, http.StatusCreated, todo.ID)

}

func TodoIsDone(w http.ResponseWriter, req *http.Request) {

	returnvars := mux.Vars(req)
	id := returnvars["id"]

	CompleteTodo(id)

	ReturnResponse(w, http.StatusAccepted, id)

}

func DeleteTodo(w http.ResponseWriter, req *http.Request) {

	returnvars := mux.Vars(req)
	id := returnvars["id"]

	RemoveTodo(id)

	ReturnResponse(w, http.StatusAccepted, id)
}

func PrepareHeaders(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Cache-control")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Accept", "application/json")

	w.WriteHeader(statusCode)
}

func ReturnResponse(w http.ResponseWriter, statusCode int, message interface{}) {
	PrepareHeaders(w, statusCode)
	json.NewEncoder(w).Encode(message)
}
