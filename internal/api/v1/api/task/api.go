package task

import (
	"github.com/go-chi/chi/v5"
	"github.com/torikki-tou/ReScheduler/internal/services/task"
	serviceDto "github.com/torikki-tou/ReScheduler/internal/services/task/dto"
	"net/http"
)

type API struct {
	service *task.Service
}

func New(service *task.Service) *API {
	return &API{
		service: service,
	}
}

func (a *API) Get(w http.ResponseWriter, r *http.Request) {
	a.service.Get(&serviceDto.GetRequest{})
	_, _ = w.Write([]byte("get"))
}

func (a *API) Search(w http.ResponseWriter, r *http.Request) {
	a.service.Search(&serviceDto.SearchRequest{})
	_, _ = w.Write([]byte("search"))
}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	a.service.Create(&serviceDto.CreateRequest{})
	_, _ = w.Write([]byte("create"))
}

func (a *API) Update(w http.ResponseWriter, r *http.Request) {
	a.service.Update(&serviceDto.UpdateRequest{})
	_, _ = w.Write([]byte("update"))
}

func (a *API) Delete(w http.ResponseWriter, r *http.Request) {
	a.service.Delete(&serviceDto.DeleteRequest{})
	_, _ = w.Write([]byte("delete"))
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
