package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"

	"github.com/dhl1402/todo-gokit/db"
	"github.com/dhl1402/todo-gokit/todo"
	"github.com/dhl1402/todo-gokit/todosvc"
)

func main() {

	mgoSession, err := mgo.Dial("localhost")
	if err != nil {
		panic(fmt.Sprintf("Cannot connect to database: %v", err))
	}

	mongo := db.NewMongo(mgoSession, "todo-gokit")
	todoRepo := todo.NewRepository(*mongo)
	todoService := todosvc.New(todoRepo)

	router := httprouter.New()
	todosvc.MakeHandler(todoService, router)

	http.ListenAndServe(":3001", router)
}
