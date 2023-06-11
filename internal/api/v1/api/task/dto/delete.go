package dto

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type (
	DeleteRequest struct {
		ID string
	}

	DeleteResponse struct {
		Task Task `json:"task"`
	}
)

func (r *DeleteRequest) Bind(req *http.Request) error {
	r.ID = chi.URLParam(req, "id")
	return nil
}
