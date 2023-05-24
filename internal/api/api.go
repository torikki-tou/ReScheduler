package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/torikki-tou/ReScheduler/internal/api/v1"
)

type API struct {
	v1 *v1.API
}

func New(v1 *v1.API) *API {
	return &API{
		v1: v1,
	}
}

func (a *API) Router() *chi.Mux {
	api := chi.NewRouter()

	api.Mount("/v1", a.v1.Router())

	router := chi.NewRouter()
	router.Mount("/api", api)
	return router
}
