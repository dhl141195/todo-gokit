package todosvc

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeGetTodosEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(GetTodosRequest)
		return svc.GetTodos(ctx, r)
	}
}

func makeCreateTodoEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(CreateTodoRequest)
		return svc.CreateTodo(ctx, r)
	}
}

func makeDeleteTodoEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(DeleteTodoRequest)
		return svc.DeleteTodo(ctx, r)
	}
}
