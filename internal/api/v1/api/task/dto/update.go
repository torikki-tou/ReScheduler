package dto

import "encoding/json"

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
