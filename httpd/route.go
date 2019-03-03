package httpd

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func (app *Application) Routes() *mux.Router {
	userChain := alice.New(metricsHandler, corsHandler, optionsHandler, jrespHandler)

	r := mux.NewRouter()

	api := r.PathPrefix(app.PrefixAPI).Subrouter()

	// Actions with auth

	api.Path("/auth/login").Handler(
		userChain.ThenFunc(app.Login),
	).Methods(http.MethodPost)

	// Actions with orders

	api.Path("/order").Handler(
		userChain.ThenFunc(app.OrdersList),
	).Methods(http.MethodGet)

	api.Path("/order").Handler(
		userChain.ThenFunc(app.CreateOrder),
	).Methods(http.MethodPost, http.MethodOptions)

	api.Path("/order/{id}").Handler(
		userChain.ThenFunc(app.OrderByID),
	).Methods(http.MethodGet)

	api.Path("/order/{id}").Handler(
		userChain.ThenFunc(app.UpdateOrder),
	).Methods(http.MethodPost)

	// Actions with clients

	api.Path("/client").Handler(
		userChain.ThenFunc(app.ClientsList),
	).Methods(http.MethodGet)

	api.Path("/client").Handler(
		userChain.ThenFunc(app.CreateClient),
	).Methods(http.MethodPost, http.MethodOptions)

	api.Path("/client/{id}").Handler(
		userChain.ThenFunc(app.ClientByID),
	).Methods(http.MethodGet)

	api.Path("/client/{id}").Handler(
		userChain.ThenFunc(app.UpdateClient),
	).Methods(http.MethodPost)

	// Actions with workers

	api.Path("/worker").Handler(
		userChain.ThenFunc(app.WorkerList),
	).Methods(http.MethodGet)

	api.Path("/worker").Handler(
		userChain.ThenFunc(app.CreateWorker),
	).Methods(http.MethodPost, http.MethodOptions)

	api.Path("/worker/{id}").Handler(
		userChain.ThenFunc(app.WorkerByID),
	).Methods(http.MethodGet)

	api.Path("/worker/{id}").Handler(
		userChain.ThenFunc(app.UpdateWorker),
	).Methods(http.MethodPost)

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return r
}
