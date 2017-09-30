package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dhl1402/todo-gokit/db"
	"github.com/dhl1402/todo-gokit/todo"
	"github.com/dhl1402/todo-gokit/todosvc"

	httptransport "github.com/go-kit/kit/transport/http"
	"gopkg.in/mgo.v2"
)

func main() {

	mgoSession, err := mgo.Dial("localhost")
	if err != nil {
		panic(fmt.Sprintf("Cannot connect to database: %v", err))
	}

	mongo := db.NewMongo(mgoSession, "todo-gokit")
	todoRepo := todo.NewRepository(*mongo)
	todoService := todosvc.New(todoRepo)

	createTodoHandler := httptransport.NewServer(
		todosvc.MakeCreateTodoEndpoint(todoService),
		decodeCreateTodoRequest,
		encodeResponse,
	)

	http.Handle("/todos/create", createTodoHandler)
	log.Fatal(http.ListenAndServe(":3001", nil))
}

func decodeCreateTodoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request todosvc.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
