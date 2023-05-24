package task

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type API struct {
}

func New() *API {
	return &API{}
}

func (a *API) Get(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(fmt.Sprint("This is a task ") + r.RequestURI))
}

func (a *API) Router() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", a.Get)

	return router
}
