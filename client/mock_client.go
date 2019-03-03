package client

import (
	"github.com/sheirys/autorepair-oms/types"
	"gopkg.in/mgo.v2/bson"
)

type MockClientService struct {
	Source map[bson.ObjectId]types.Client
}

func (m *MockClientService) Init() error {
	m.Source = make(map[bson.ObjectId]types.Client)
	return nil
}

func (m *MockClientService) Get(id bson.ObjectId) (types.Client, error) {
	if client, ok := m.Source[id]; ok {
		return client, nil
	}
	return types.Client{}, types.ErrDocumentNotFound
}

func (m *MockClientService) All() ([]types.Client, error) {
	list := []types.Client{}
	for _, v := range m.Source {
		list = append(list, v)
	}
	return list, nil
}

func (m *MockClientService) Create(c types.Client) (types.Client, error) {
	c.ID = bson.NewObjectId()
	m.Source[c.ID] = c
	return c, nil
}

func (m *MockClientService) Update(c types.Client) (types.Client, error) {
	if _, ok := m.Source[c.ID]; ok {
		m.Source[c.ID] = c
		return c, nil
	}
	return types.Client{}, types.ErrDocumentNotFound
}
