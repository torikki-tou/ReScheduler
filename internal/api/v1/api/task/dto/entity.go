package dto

import "encoding/json"

type Task struct {
	ID             string          `json:"id"`
	CronExpression string          `json:"cron"`
	Message        json.RawMessage `json:"message"`
}
