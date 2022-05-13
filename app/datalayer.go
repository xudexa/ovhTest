package app

func ListTodos(params QueryParams) []Todo {

	var (
		todos []Todo
	)

	s := GetServer()

	switch s.DbType {
	case "postgresql":

	default:
		todos = s.TodosStore.ListTodosInMemory()
	}

	return todos
}

func AddTodo(t Todo) {
	s := GetServer()

	switch s.DbType {
	case "postgresql":
		AddTodoInPostgresql(t)
	default:
		s.TodosStore.AddTodoInMemory(t)
	}
}

func CompleteTodo(id string) {
	s := GetServer()

	switch s.DbType {
	case "postgresql":
		CompleteTodoInPostgresql(id)
	default:
		s.TodosStore.CompleteTodoInMemory(id)
	}
}

func RemoveTodo(id string) {
	s := GetServer()

	switch s.DbType {
	case "postgresql":
		RemoveTodoInPostgresql(id)
	default:
		s.TodosStore.RemoveTodoInMemory(id)
	}
}
