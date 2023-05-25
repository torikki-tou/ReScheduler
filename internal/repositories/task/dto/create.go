package dto

import "encoding/json"

type (
	CreateRequest struct {
		ID             string
		CronExpression string
		Message        json.RawMessage
	}

	CreateResponse struct {
		Task Task
	}
)
