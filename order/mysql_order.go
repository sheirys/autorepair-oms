package order

import (
	"github.com/sheirys/autorepair-oms/types"
	"gopkg.in/mgo.v2/bson"
)

type MySQLOrderService struct {
}

func (s *MySQLOrderService) Init() error {
	return nil
}

func (s *MySQLOrderService) Get(id bson.ObjectId) (types.Order, error) {
	var order types.Order
	return order, nil
}

func (s *MySQLOrderService) All() ([]types.Order, error) {
	list := []types.Order{}
	return list, nil
}

func (s *MySQLOrderService) Create(o types.Order) (types.Order, error) {
	return s.Get(o.ID)
}

func (s *MySQLOrderService) Update(o types.Order) (types.Order, error) {
	return s.Get(o.ID)
}
