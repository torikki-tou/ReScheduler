package dto

import (
	"encoding/json"
	"net/http"
)

type (
	CreateRequest struct {
		ID             *string         `json:"id"`
		CronExpression string          `json:"cron"`
		Message        json.RawMessage `json:"message"`
	}

	CreateResponse struct {
		Task Task `json:"task"`
	}
)

func (r *CreateRequest) Bind(req *http.Request) error {
	return json.NewDecoder(req.Body).Decode(&r)
}
