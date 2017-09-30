package todosvc

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeCreateTodoEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(CreateTodoRequest)
		return svc.CreateTodo(ctx, r)
	}
}
