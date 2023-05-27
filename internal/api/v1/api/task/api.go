package task

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
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

	router.Get("/", a.Search)
	router.Get("/{id}", a.Get)
	router.Post("/", a.Create)
	router.Patch("/{id}", a.Update)
	router.Delete("/{id}", a.Delete)

	return router
}

func (a *API) Get(w http.ResponseWriter, r *http.Request) {

	var req dto.GetRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return
	}

	res := a.service.Get(&serviceDto.GetRequest{ID: req.ID})

	err = json.NewEncoder(w).Encode(dto.GetResponse{Task: *a.fromService(&res.Task)})
	if err != nil {
		return
	}
}

func (a *API) Search(w http.ResponseWriter, r *http.Request) {

	var req dto.SearchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return
	}

	res := a.service.Search(&serviceDto.SearchRequest{Limit: req.Limit})

	var tasks = make([]dto.Task, 0, len(res.Tasks))
	for _, task := range res.Tasks {
		tasks = append(tasks, *a.fromService(&task))
	}

	err = json.NewEncoder(w).Encode(dto.SearchResponse{Tasks: tasks})
	if err != nil {
		return
	}
}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {

	var req dto.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return
	}

	res := a.service.Create(&serviceDto.CreateRequest{
		ID:             req.ID,
		CronExpression: req.CronExpression,
		Message:        req.Message,
	})

	err = json.NewEncoder(w).Encode(dto.CreateResponse{Task: *a.fromService(&res.Task)})
	if err != nil {
		return
	}
}

func (a *API) Update(w http.ResponseWriter, r *http.Request) {

	var req dto.UpdateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return
	}

	res := a.service.Update(&serviceDto.UpdateRequest{
		ID: req.ID,
		Update: serviceDto.TaskUpdate{
			ID:             req.Update.ID,
			CronExpression: req.Update.CronExpression,
			Message:        req.Update.Message,
		},
	})

	err = json.NewEncoder(w).Encode(dto.UpdateResponse{Task: *a.fromService(&res.Task)})
	if err != nil {
		return
	}
}

func (a *API) Delete(w http.ResponseWriter, r *http.Request) {

	var req dto.DeleteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return
	}

	res := a.service.Delete(&serviceDto.DeleteRequest{ID: req.ID})

	err = json.NewEncoder(w).Encode(dto.DeleteResponse{Task: *a.fromService(&res.Task)})
	if err != nil {
		return
	}
}

func (a *API) fromService(req *serviceDto.Task) *dto.Task {
	return &dto.Task{
		ID:             req.ID,
		CronExpression: req.CronExpression,
		Message:        req.Message,
	}
}
