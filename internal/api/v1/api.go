package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/torikki-tou/ReScheduler/internal/api/v1/api/task"
)

type API struct {
	task *task.API
}

func New(task *task.API) *API {
	return &API{
		task: task,
	}
}

func (a *API) Router() *chi.Mux {
	router := chi.NewRouter()

	router.Mount("/tasks", a.task.Router())

	return router
}
