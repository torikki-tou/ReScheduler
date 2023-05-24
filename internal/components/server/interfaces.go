package server

import "net/http"

type (
	mux interface {
		ServeHTTP(http.ResponseWriter, *http.Request)
	}
	config interface {
		AppPort() int
	}
)
