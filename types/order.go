package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Order defines what will be saved into database.
type Order struct {
	ID     bson.ObjectId `json:"order_id" bson:"_id"`
	Meta   Meta          `json:"meta" bson:"meta"`
	Client Client        `json:"client" bson:"client"`
	Car    Car           `json:"car" bson:"car"`
	Todo   []Todo        `json:"todo" bson:"todo"`
	Job    []Job         `json:"job" bson:"job"`
	Part   []Part        `json:"part" bson:"part"`
}

type Meta struct {
	WorkerName string `json:"worker_name" bson:"worker_name"`
	//	WorkerID      bson.ObjectId `json:"worker_id" bson:"worker_id"`
	AcceptTime    time.Time `json:"accept_time" bson:"accept_time"`
	PaymentMethod string    `json:"payment_method" bson:"payment_method"`
	PaymentStatus string    `json:"payment_status" bson:"payment_status"`
	Invoice       string    `json:"invoice" bson:"invoice"`
	Description   string    `json:"description" bson:"description"`
	Status        string    `json:"status" bson:"status"`
}

type Todo struct {
	Description string `json:"description"`
}

type Job struct {
	Description string  `json:"description" bson:"description"`
	WorkerName  string  `json:"worker_name" bson:"worker_name"`
	Quantity    uint32  `json:"quantity" bson:"quantity"`
	Price       float32 `json:"price" bson:"price"`
	Done        bool    `json:"done" bson:"done"`
}

type Part struct {
	Name     string  `json:"name" bson:"name"`
	Code     string  `json:"code" bson:"code"`
	Quantity uint32  `json:"quantity" bson:"quantity"`
	Price    float32 `json:"price" bson:"price"`
}

// OrderService defines what to expect from order service.
type OrderService interface {
	Init() error
	Get(id bson.ObjectId) (Order, error)
	All() ([]Order, error)
	Create(o Order) (Order, error)
	Update(Order) (Order, error)
}
