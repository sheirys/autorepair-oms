package order

import (
	"github.com/sheirys/autorepair-oms/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoOrderService struct {
	DB *mgo.Collection
}

func (os *MongoOrderService) Init() error {
	return nil
}

func (os *MongoOrderService) Get(id bson.ObjectId) (types.Order, error) {
	var order types.Order
	if err := os.DB.FindId(id).One(&order); err != nil {
		return order, types.ErrDocumentNotFound
	}
	return order, nil
}

func (os *MongoOrderService) All() ([]types.Order, error) {
	list := []types.Order{}
	if err := os.DB.Find(bson.M{}).Sort("-meta.accept_time").All(&list); err != nil {
		return nil, err
	}
	return list, nil
}

func (os *MongoOrderService) Create(o types.Order) (types.Order, error) {
	o.ID = bson.NewObjectId()
	if err := os.DB.Insert(o); err != nil {
		return o, err
	}
	return os.Get(o.ID)
}

func (os *MongoOrderService) Update(o types.Order) (types.Order, error) {
	if _, err := os.Get(o.ID); err == types.ErrDocumentNotFound {
		return types.Order{}, types.ErrDocumentNotFound
	}
	if err := os.DB.UpdateId(o.ID, o); err != nil {
		return types.Order{}, err
	}
	return os.Get(o.ID)
}
