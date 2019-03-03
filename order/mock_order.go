package order

import (
	"github.com/sheirys/autorepair-oms/types"
	"gopkg.in/mgo.v2/bson"
)

type MockOrderService struct {
	Source map[bson.ObjectId]types.Order
}

func (m *MockOrderService) Init() error {
	m.Source = make(map[bson.ObjectId]types.Order)
	return nil
}

func (m *MockOrderService) Get(id bson.ObjectId) (types.Order, error) {
	if order, ok := m.Source[id]; ok {
		return order, nil
	}
	return types.Order{}, types.ErrDocumentNotFound
}

func (m *MockOrderService) All() ([]types.Order, error) {
	list := []types.Order{}
	for _, v := range m.Source {
		list = append(list, v)
	}
	return list, nil
}

func (m *MockOrderService) Create(o types.Order) (types.Order, error) {
	o.ID = bson.NewObjectId()
	m.Source[o.ID] = o
	return o, nil
}

func (m *MockOrderService) Update(o types.Order) (types.Order, error) {
	if _, ok := m.Source[o.ID]; ok {
		m.Source[o.ID] = o
		return o, nil
	}
	return types.Order{}, types.ErrDocumentNotFound
}
