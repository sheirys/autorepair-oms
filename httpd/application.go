package httpd

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/sheirys/autorepair-oms/types"
	"github.com/sirupsen/logrus"
)

type Application struct {
	Version   Version
	Listen    string
	PrefixAPI string
	Backend   Backend
	Debug     bool
	httpd     http.Server

	// SignSecret is used to sign JWT tokens. New sign secret should be
	// generated everytime on startup.
	SignSecret []byte
}

func (app *Application) Start() error {

	// generate secret key for jwt sessions.
	h := sha256.New()
	h.Write([]byte(time.Now().String()))
	app.SignSecret = []byte(fmt.Sprintf("%x", h.Sum(nil)))

	// export secret key to tmp file so it can be accessible for debugging.
	tmpFile, err := ioutil.TempFile(os.TempDir(), fmt.Sprintf("axisd_%d_sign", os.Getpid()))
	if err != nil {
		logrus.WithError(err).Error("cannot create temporary file")
		return err
	}
	if _, err = tmpFile.Write(app.SignSecret); err != nil {
		logrus.WithError(err).Error("cannot export sign secret")
		return err
	}

	// sanitize config and set default values if nothis was set.
	if app.PrefixAPI == "" {
		app.PrefixAPI = "/api/v1/"
	}
	if app.Listen == "" {
		app.Listen = ":8081"
	}

	// set routes to handle for this application.
	app.httpd.Handler = app.Routes()

	if err := app.Backend.OrderService.Init(); err != nil {
		return err
	}
	if err := app.Backend.ClientService.Init(); err != nil {
		return err
	}
	if err := app.Backend.WorkerService.Init(); err != nil {
		return err
	}

	logrus.Infof("Starting application on %s", app.Listen)
	app.httpd.Addr = app.Listen
	return app.httpd.ListenAndServe()
}

type Backend struct {
	OrderService  types.OrderService
	ClientService types.ClientService
	WorkerService types.WorkerService
}
