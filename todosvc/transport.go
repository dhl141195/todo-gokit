package todosvc

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"
)

func MakeHandler(s Service, router *httprouter.Router) {
	router.GET("/todos", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		kithttp.NewServer(
			makeGetTodosEndpoint(s),
			decodeGetTodosRequest,
			kithttp.EncodeJSONResponse,
		).ServeHTTP(w, r)
	})

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

func decodeGetTodosRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	q := r.URL.Query()
	request := GetTodosRequest{}

	if v := q.Get("status"); v != "" {
		request.Status = v
	}

	if v := q.Get("keyword"); v != "" {
		request.Keyword = v
	}

	if v := q.Get("limit"); v != "" {
		request.Limit, _ = strconv.Atoi(v)
	}

	return request, nil
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
	request := DeleteTodoRequest{
		ID: ps.ByName("todoId"),
	}
	return request, nil
}
