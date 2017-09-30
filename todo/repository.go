package todo

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/dhl1402/todo-gokit/db"
)

const collectionName = "todos"

// Errors
var (
	ErrTodoNotExist = errors.New("Todo not exist")
)

// Repository is interface to access persistence layer
type Repository interface {
	GetByID(id string) (*Todo, error)
	Save(todo *Todo) error
	Delete(id string) error
	// Get(*db.Query) ([]Todo, *db.Paging, error)
}

type repository struct {
	mongo db.Mongo
}

func NewRepository(mongo db.Mongo) Repository {
	return &repository{
		mongo: mongo,
	}
}

func (repo *repository) GetByID(id string) (*Todo, error) {
	var todo Todo
	err := repo.mongo.DB.C(collectionName).FindId(id).One(&todo)
	if err == mgo.ErrNotFound {
		return nil, ErrTodoNotExist
	} else if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (repo *repository) Save(todo *Todo) error {
	if todo.ID == "" {
		todo.ID = bson.NewObjectId().Hex()
	}
	_, err := repo.mongo.DB.C(collectionName).UpsertId(todo.ID, todo)

	return err
}

func (repo *repository) Delete(id string) error {
	err := repo.mongo.DB.C(collectionName).RemoveId(id)
	return err
}

// TODO: Get
