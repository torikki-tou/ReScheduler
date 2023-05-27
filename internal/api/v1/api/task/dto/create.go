package dto

import "encoding/json"

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
