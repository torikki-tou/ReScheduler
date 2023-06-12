package dto

import "encoding/json"

type (
	UpdateRequest struct {
		ID     string
		Update TaskUpdate
	}
	TaskUpdate struct {
		ID             *string
		CronExpression *string
		Message        *json.RawMessage
	}
)
