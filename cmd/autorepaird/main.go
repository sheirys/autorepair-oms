package main

import (
	"time"

	"github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"

	"github.com/sheirys/autorepair-oms/client"
	"github.com/sheirys/autorepair-oms/httpd"
	"github.com/sheirys/autorepair-oms/order"
	"github.com/sheirys/autorepair-oms/worker"
)

var (
	Commit    string
	Branch    string
	BuildDate string
	Tag       string
)

func main() {
	// connect to mongo database
	session, err := mgo.DialWithTimeout("localhost", time.Second*5)
	if err != nil {
		logrus.WithError(err).Fatal("cannot connect to mongo")
		return
	}

	app := &httpd.Application{
		Debug: true,
		Backend: httpd.Backend{
			OrderService:  &order.MongoOrderService{DB: session.DB("axis").C("orders")},
			ClientService: &client.MongoClientService{DB: session.DB("axis").C("clients")},
			WorkerService: &worker.MongoWorkerService{DB: session.DB("axis").C("workers")},
		},
		Version: httpd.Version{
			Commit:    Commit,
			Branch:    Branch,
			BuildDate: BuildDate,
			Version:   Tag,
		},
	}

	if err := app.Start(); err != nil {
		logrus.WithError(err).Error("cannot start application")
	}

}
