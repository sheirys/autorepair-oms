package httpd

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sheirys/autorepair-oms/types"
	"github.com/sheirys/autorepair-oms/utils/api"
	"github.com/sirupsen/logrus"
)

// OrdersList will extract all available orders for this user.
// Endpoint: [GET] /api/v1/order
func (app *Application) OrdersList(w http.ResponseWriter, r *http.Request) {
	list, err := app.Backend.OrderService.All()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.WithError(err).Error("cannot extract orders")
		return
	}
	api.JSON(w, http.StatusOK, list)
}

// CreateOrder will create new order
// Endpoint: [POST] /api/v1/order
func (app *Application) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var (
		order  types.Order
		client types.Client
		err    error
	)
	if ok, err := api.BindJSON(r, &order); !ok {
		logrus.WithError(err).Error("parsing request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// create client if it is new client
	client, err = app.Backend.ClientService.Get(order.Client.ID)
	if err == types.ErrDocumentNotFound {
		// create new client
		client, err = app.Backend.ClientService.Create(order.Client)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logrus.WithError(err).Error("cannot create client")
			return
		}
	}
	order.Client = client
	order.Meta.AcceptTime = time.Now()
	if order, err = app.Backend.OrderService.Create(order); err != nil {
		logrus.WithError(err).Error("creating order")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	order, err = app.Backend.OrderService.Get(order.ID)
	if err != nil {
		logrus.WithError(err).Error("cannot extract created order")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	api.JSON(w, http.StatusOK, order)
}

// OrderByID will extract single order
// Endpoint: [GET] /api/v1/order/{id}
func (app *Application) OrderByID(w http.ResponseWriter, r *http.Request) {
	id := api.SegmentBsonID(mux.Vars(r), "id")
	order, err := app.Backend.OrderService.Get(id)
	if err == types.ErrDocumentNotFound {
		logrus.Warnf("order '%s' not found", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		logrus.WithError(err).Errorf("extracting order '%s'", id)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	api.JSON(w, http.StatusOK, order)
}

// UpdateOrder will update single order
// Endpoint: [POST] /api/v1/order/{id}
func (app *Application) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	id := api.SegmentBsonID(mux.Vars(r), "id")
	order, err := app.Backend.OrderService.Get(id)
	if err == types.ErrDocumentNotFound {
		logrus.Warnf("order '%s' not found", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if ok, err := api.BindJSON(r, &order); !ok {
		logrus.WithError(err).Error("parsing request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	order.ID = id
	order, err = app.Backend.OrderService.Update(order)
	if err != nil {
		logrus.WithError(err).Error("updating order")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	api.JSON(w, http.StatusOK, order)
}

// ClientsList will extract all available clients for this user.
// Endpoint: [GET] /api/v1/client
func (app *Application) ClientsList(w http.ResponseWriter, r *http.Request) {
	list, err := app.Backend.ClientService.All()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.WithError(err).Error("cannot extract orders")
		return
	}
	api.JSON(w, http.StatusOK, list)
}

// ClientByID will extract single client
// Endpoint: [GET] /api/v1/client/{id}
func (app *Application) ClientByID(w http.ResponseWriter, r *http.Request) {
	id := api.SegmentBsonID(mux.Vars(r), "id")
	client, err := app.Backend.ClientService.Get(id)
	if err == types.ErrDocumentNotFound {
		w.WriteHeader(http.StatusNotFound)
		logrus.WithError(err).Warn("cannot extract client")
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.WithError(err).Error("cannot extract client")
		return
	}
	api.JSON(w, http.StatusOK, client)
}

// CreateClient will create new client
// Endpoint: [POST] /api/v1/client
func (app *Application) CreateClient(w http.ResponseWriter, r *http.Request) {
	var (
		client types.Client
		err    error
	)
	if ok, err := api.BindJSON(r, &client); !ok {
		logrus.WithError(err).Error("parsing request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if client, err = app.Backend.ClientService.Create(client); err != nil {
		logrus.WithError(err).Error("creating client")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	client, err = app.Backend.ClientService.Get(client.ID)
	if err != nil {
		logrus.WithError(err).Error("cannot extract created client")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	api.JSON(w, http.StatusOK, client)
}

// UpdateClient will update client
// Endpoint: [POST] /api/v1/client/{id}
func (app *Application) UpdateClient(w http.ResponseWriter, r *http.Request) {
	id := api.SegmentBsonID(mux.Vars(r), "id")
	client, err := app.Backend.ClientService.Get(id)
	if err == types.ErrDocumentNotFound {
		logrus.WithError(err).Warnf("client '%s' not found", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if ok, err := api.BindJSON(r, &client); !ok {
		logrus.WithError(err).Error("parsing request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	client, err = app.Backend.ClientService.Update(client)
	if err != nil {
		logrus.WithError(err).Error("updating client")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	api.JSON(w, http.StatusOK, client)
}

// WorkerList will return all available workers
// Endpoint: [GET] /api/v1/worker
func (app *Application) WorkerList(w http.ResponseWriter, r *http.Request) {
	list, err := app.Backend.WorkerService.All()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.WithError(err).Error("cannot extract orders")
		return
	}
	api.JSON(w, http.StatusOK, list)
}

// CreateWorker will create new worker
// Endpoint: [POST] /api/v1/worker
func (app *Application) CreateWorker(w http.ResponseWriter, r *http.Request) {
	var (
		worker types.Worker
		err    error
	)
	if ok, err := api.BindJSON(r, &worker); !ok {
		logrus.WithError(err).Error("parsing request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if worker, err = app.Backend.WorkerService.Create(worker); err != nil {
		logrus.WithError(err).Error("creating worker")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	worker, err = app.Backend.WorkerService.Get(worker.ID)
	if err != nil {
		logrus.WithError(err).Error("cannot extract created worker")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	api.JSON(w, http.StatusOK, worker)
}

// WorkerByID will extract single worker by provided id
// Endpoint: [GET] /api/v1/worker/{id}
func (app *Application) WorkerByID(w http.ResponseWriter, r *http.Request) {
	id := api.SegmentBsonID(mux.Vars(r), "id")
	worker, err := app.Backend.WorkerService.Get(id)
	if err == types.ErrDocumentNotFound {
		w.WriteHeader(http.StatusNotFound)
		logrus.WithError(err).Warn("cannot extract worker")
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.WithError(err).Error("cannot extract worker")
		return
	}
	api.JSON(w, http.StatusOK, worker)
}

// UpdateWorker will update single worker.
// Endpoint: [POST] /api/v1/worker/{id}
func (app *Application) UpdateWorker(w http.ResponseWriter, r *http.Request) {
	var (
		worker types.Worker
		err    error
	)
	if ok, err := api.BindJSON(r, &worker); !ok {
		logrus.WithError(err).Error("parsing request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if worker, err = app.Backend.WorkerService.Create(worker); err != nil {
		logrus.WithError(err).Error("creating worker")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	worker, err = app.Backend.WorkerService.Get(worker.ID)
	if err != nil {
		logrus.WithError(err).Error("cannot extract created client")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	api.JSON(w, http.StatusOK, worker)
}
