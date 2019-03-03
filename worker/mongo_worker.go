package worker

import (
	"github.com/sheirys/autorepair-oms/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoWorkerService struct {
	DB *mgo.Collection
}

func (os *MongoWorkerService) Init() error {
	return nil
}

func (os *MongoWorkerService) Get(id bson.ObjectId) (types.Worker, error) {
	var worker types.Worker
	if err := os.DB.FindId(id).One(&worker); err != nil {
		return worker, types.ErrDocumentNotFound
	}
	return worker, nil
}

func (os *MongoWorkerService) All() ([]types.Worker, error) {
	list := []types.Worker{}
	if err := os.DB.Find(bson.M{}).All(&list); err != nil {
		return nil, err
	}
	return list, nil
}

func (os *MongoWorkerService) Create(o types.Worker) (types.Worker, error) {
	o.ID = bson.NewObjectId()
	if err := os.DB.Insert(o); err != nil {
		return o, err
	}
	return os.Get(o.ID)
}

func (os *MongoWorkerService) Update(o types.Worker) (types.Worker, error) {
	if _, err := os.Get(o.ID); err == types.ErrDocumentNotFound {
		return types.Worker{}, types.ErrDocumentNotFound
	}
	if err := os.DB.UpdateId(o.ID, o); err != nil {
		return types.Worker{}, err
	}
	return os.Get(o.ID)
}
