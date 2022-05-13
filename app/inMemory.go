package app

import (
	"ovhTest/app/functions"
	"time"
)

// Encapsulation
type TodosStore struct {
	todos map[string]Todo
}

func NewTodoStore() TodosStore {
	return TodosStore{make(map[string]Todo)}
}

func (ts *TodosStore) ListTodosInMemory() []Todo {
	var (
		todos []Todo
	)

	for _, todo := range ts.todos {
		todos = append(todos, todo)
	}

	return todos
}

func (ts *TodosStore) AddTodoInMemory(t Todo) {
	t.ID = functions.NewUUID()
	t.CreatedAt = time.Now()
	t.Completed = false
	ts.todos[t.ID] = t
}

func (ts *TodosStore) CompleteTodoInMemory(id string) {
	todo := ts.todos[id]
	todo.Completed = !todo.Completed
	todo.CompletedAt = time.Now()
	ts.todos[id] = todo

}

func (ts *TodosStore) RemoveTodoInMemory(id string) {
	delete(ts.todos, id)
}
