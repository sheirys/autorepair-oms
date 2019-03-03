package httpd

import (
	"net/http"
	"time"
)

type Version struct {
	Commit     string
	Branch     string
	BuildDate  string
	Version    string
	UptimeFrom time.Time
}

func (app *Application) APIVersion(w http.ResponseWriter, r *http.Request) {
}
