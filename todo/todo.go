package todo

import (
	"time"
)

type Todo struct {
	ID         string    `bson:"_id" json:"id"`
	Name       string    `json:"name"`
	Status     bool      `json:"status"`
	CreateAt   time.Time `json:"createAt"`
	LastUpdate time.Time `json:"lastUpdate"`
}
