package client

import (
	"github.com/sheirys/autorepair-oms/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoClientService struct {
	DB *mgo.Collection
}

func (c *MongoClientService) Init() error {
	return nil
}

func (c *MongoClientService) Get(id bson.ObjectId) (types.Client, error) {
	client := types.Client{}
	if err := c.DB.FindId(id).One(&client); err != nil {
		return client, types.ErrDocumentNotFound
	}
	return client, nil
}

func (c *MongoClientService) All() ([]types.Client, error) {
	list := []types.Client{}
	if err := c.DB.Find(bson.M{}).All(&list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *MongoClientService) Create(client types.Client) (types.Client, error) {
	client.ID = bson.NewObjectId()
	if err := c.DB.Insert(client); err != nil {
		return client, err
	}
	return c.Get(client.ID)
}

func (c *MongoClientService) Update(client types.Client) (types.Client, error) {
	if _, err := c.Get(client.ID); err == types.ErrDocumentNotFound {
		return types.Client{}, types.ErrDocumentNotFound
	}
	if err := c.DB.UpdateId(client.ID, client); err != nil {
		return types.Client{}, err
	}
	return c.Get(client.ID)
}
