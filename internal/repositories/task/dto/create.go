package dto

import "encoding/json"

type (
	CreateRequest struct {
		ID             string
		Score          int64
		CronExpression string
		Message        json.RawMessage
	}

	CreateResponse struct {
		Task Task
	}
)
