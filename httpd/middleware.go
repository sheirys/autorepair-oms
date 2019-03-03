package httpd

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// metricsHandler is used to measure request proccess time and its lifetime.
// function should be chained within alice chain builder.
func metricsHandler(next http.Handler) http.Handler {
	mw := func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Now().Sub(startTime)
		logrus.WithFields(logrus.Fields{
			"duration": duration,
			"url":      r.URL.Path,
			"method":   r.Method,
		}).Info("metrics")
	}
	return http.HandlerFunc(mw)
}

// enable cors for this handler
func corsHandler(next http.Handler) http.Handler {
	mw := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(mw)
}

// enable cors for this handler
func optionsHandler(next http.Handler) http.Handler {
	mw := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodOptions {
			next.ServeHTTP(w, r)
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}
	return http.HandlerFunc(mw)
}

// jrespHandler is used to decorate response content type to application/json.
func jrespHandler(next http.Handler) http.Handler {
	mw := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(mw)
}
