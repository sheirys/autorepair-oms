package types

import (
	"gopkg.in/mgo.v2/bson"
)

type Worker struct {
	ID   bson.ObjectId `json:"worker_id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}

// OrderService defines what to expect from order service.
type WorkerService interface {
	Init() error
	Get(id bson.ObjectId) (Worker, error)
	All() ([]Worker, error)
	Create(o Worker) (Worker, error)
	Update(Worker) (Worker, error)
}
