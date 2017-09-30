package todosvc

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/dhl1402/todo-gokit/todo"
)

// Service is interface of user apis service
type Service interface {
	// GetTodos(ctx context.Context, r GetTodosRequest) (*GetTodosResponse, error)
	CreateTodo(ctx context.Context, r CreateTodoRequest) (*CreateTodoResponse, error)
	// UpdateTodo(ctx context.Context, r UpdateTodoRequest) (*UpdateTodoResponse, error)
	DeleteTodo(ctx context.Context, r DeleteTodoRequest) (*DeleteTodoResponse, error)
}

type service struct {
	todoRepo todo.Repository
}

func New(todoRepo todo.Repository) Service {
	return &service{
		todoRepo: todoRepo,
	}
}

func (s *service) CreateTodo(ctx context.Context, r CreateTodoRequest) (*CreateTodoResponse, error) {
	r.Name = strings.TrimSpace(r.Name)
	if r.Name == "" {
		return nil, errors.New("Missing name")
	}

	now := time.Now().UTC()
	todo := &todo.Todo{
		Name:       r.Name,
		Status:     "pending",
		CreateAt:   now,
		LastUpdate: now,
	}

	err := s.todoRepo.Save(todo)
	if err != nil {
		return nil, err
	}

	return &CreateTodoResponse{
		TodoResponse: getTodoResponse(todo),
	}, nil
}

func (s *service) DeleteTodo(ctx context.Context, r DeleteTodoRequest) (*DeleteTodoResponse, error) {
	err := s.todoRepo.Delete(r.ID)
	if err != nil {
		return &DeleteTodoResponse{
			Status: "Error",
			Error:  err.Error(),
		}, err
	}

	return &DeleteTodoResponse{
		Status: "Success",
		Error:  "",
	}, nil
}

func getTodoResponse(todo *todo.Todo) TodoResponse {
	return TodoResponse{
		Todo: todo,
	}
}
