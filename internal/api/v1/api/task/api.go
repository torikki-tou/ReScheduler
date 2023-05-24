package task

import (
	"github.com/go-chi/chi/v5"
	"github.com/torikki-tou/ReScheduler/internal/services/task"
	"net/http"
)

type API struct {
	service *task.Service
}

func New() *API {
	return &API{}
}

func (a *API) Get(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(a.service.Get() + " " + chi.URLParam(r, "id")))
}

func (a *API) Search(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(a.service.Search()))
}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(a.service.Create()))
}

func (a *API) Update(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(a.service.Update() + " " + chi.URLParam(r, "id")))
}

func (a *API) Delete(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(a.service.Delete() + " " + chi.URLParam(r, "id")))
}
func (a *API) Router() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", a.Search)
	router.Get("/{id}", a.Get)
	router.Post("/", a.Create)
	router.Patch("/{id}", a.Update)
	router.Delete("/{id}", a.Delete)

	return router
}
