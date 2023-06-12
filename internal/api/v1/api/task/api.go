package task

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/torikki-tou/ReScheduler/internal/api/v1/api/task/dto"
	taskService "github.com/torikki-tou/ReScheduler/internal/services/task"
	serviceDto "github.com/torikki-tou/ReScheduler/internal/services/task/dto"
	"net/http"
)

type API struct {
	service *taskService.Service
}

func New(service *taskService.Service) *API {
	return &API{
		service: service,
	}
}

func (a *API) Router() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{id}", a.Get)
	router.Post("/", a.Create)
	router.Patch("/{id}", a.Update)
	router.Delete("/{id}", a.Delete)

	return router
}

func (a *API) Get(w http.ResponseWriter, r *http.Request) {

	var req dto.GetRequest
	err := req.Bind(r)
	if err != nil {
		return
	}

	res, err := a.service.Get(&serviceDto.GetRequest{ID: req.ID})
	if err != nil {
		return
	}

	render.JSON(w, r, res)
}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {

	var req dto.CreateRequest
	err := req.Bind(r)
	if err != nil {
		return
	}

	res, err := a.service.Create(&serviceDto.CreateRequest{
		ID:             req.ID,
		CronExpression: req.CronExpression,
		Message:        req.Message,
	})
	if err != nil {
		return
	}

	render.JSON(w, r, dto.CreateResponse{Task: *a.fromService(&res.Task)})
}

func (a *API) Update(w http.ResponseWriter, r *http.Request) {

	var req dto.UpdateRequest
	err := req.Bind(r)
	if err != nil {
		return
	}

	res, err := a.service.Update(&serviceDto.UpdateRequest{
		ID: req.ID,
		Update: serviceDto.TaskUpdate{
			ID:             req.Update.ID,
			CronExpression: req.Update.CronExpression,
			Message:        req.Update.Message,
		},
	})
	if err != nil {
		return
	}

	render.JSON(w, r, dto.UpdateResponse{Task: *a.fromService(&res.Task)})
}

func (a *API) Delete(w http.ResponseWriter, r *http.Request) {

	var req dto.DeleteRequest
	err := req.Bind(r)
	if err != nil {
		return
	}

	res, err := a.service.Delete(&serviceDto.DeleteRequest{ID: req.ID})
	if err != nil {
		return
	}

	render.JSON(w, r, dto.DeleteResponse{Task: *a.fromService(&res.Task)})
}

func (a *API) fromService(req *serviceDto.Task) *dto.Task {
	return &dto.Task{
		ID:             req.ID,
		CronExpression: req.CronExpression,
		Message:        req.Message,
	}
}
