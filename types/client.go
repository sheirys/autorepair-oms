package types

import "gopkg.in/mgo.v2/bson"

type Client struct {
	ID          bson.ObjectId   `json:"client_id" bson:"_id"`
	Name        string          `json:"name" bson:"name"`
	Email       string          `json:"email" bson:"email"`
	Phone       string          `json:"phone" bson:"phone"`
	Description string          `json:"description" bson:"description"`
	Orders      []bson.ObjectId `json:"orders" bson:"orders"`
	Cars        []bson.ObjectId `json:"cars" bson:"cars"`
}

// ClientService defines what to expect from order service.
type ClientService interface {
	Init() error
	Get(id bson.ObjectId) (Client, error)
	All() ([]Client, error)
	Create(c Client) (Client, error)
	Update(Client) (Client, error)
}
