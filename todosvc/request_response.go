package todosvc

import "github.com/dhl1402/todo-gokit/todo"

type TodoResponse struct {
	*todo.Todo
}

type GetTodosRequest struct {
	Limit int
}

type GetTodosResponse struct {
	Total int            `json:"total"`
	Todos []TodoResponse `json:"todos"`
}

type CreateTodoRequest struct {
	Name string `json:"name"`
}

type CreateTodoResponse struct {
	TodoResponse
}

type UpdateTodoRequest struct {
	ID string
	CreateTodoRequest
	Status bool `json:"status"`
}

type UpdateTodoResponse struct {
	TodoResponse
}

type DeleteTodoRequest struct {
	ID string
}

type DeleteTodoResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}
