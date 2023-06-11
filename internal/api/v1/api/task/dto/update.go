package dto

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type (
	UpdateRequest struct {
		ID     string     `json:"id"`
		Update TaskUpdate `json:"update"`
	}
	TaskUpdate struct {
		ID             *string          `json:"id"`
		CronExpression *string          `json:"cron"`
		Message        *json.RawMessage `json:"message"`
	}

	UpdateResponse struct {
		Task Task
	}
)

func (r *UpdateRequest) Bind(req *http.Request) error {
	r.ID = chi.URLParam(req, "id")

	var update TaskUpdate
	err := json.NewDecoder(req.Body).Decode(&update)
	if err != nil {
		return err
	}

	r.Update = update
	return nil
}
