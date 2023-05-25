package dto

import "encoding/json"

type Task struct {
	ID             string
	CronExpression string
	Message        json.RawMessage
}
