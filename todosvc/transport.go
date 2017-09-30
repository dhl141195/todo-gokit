package todosvc

import (
	"context"
	"encoding/json"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"
)

func MakeHandler(s Service, router *httprouter.Router) {
	router.POST("/todos", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		kithttp.NewServer(
			makeCreateTodoEndpoint(s),
			decodeCreateTodoRequest,
			kithttp.EncodeJSONResponse,
		).ServeHTTP(w, r)
	})

	router.DELETE("/todos/:todoId", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handler := kithttp.NewServer(
			makeDeleteTodoEndpoint(s),
			decodeDeleteTodoRequest,
			kithttp.EncodeJSONResponse,
		)

		r = r.WithContext(context.WithValue(r.Context(), "URIParams", ps))
		handler.ServeHTTP(w, r)
	})
}

func decodeCreateTodoRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeDeleteTodoRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	ps := ctx.Value("URIParams").(httprouter.Params)
	req := DeleteTodoRequest{
		ID: ps.ByName("todoId"),
	}
	return req, nil
}
