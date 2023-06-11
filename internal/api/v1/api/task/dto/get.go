package dto

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type (
	GetRequest struct {
		ID string
	}

	GetResponse struct {
		Task Task `json:"task"`
	}
)

func (r *GetRequest) Bind(req *http.Request) error {
	r.ID = chi.URLParam(req, "id")
	return nil
}
