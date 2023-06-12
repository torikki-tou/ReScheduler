package dto

import "encoding/json"

type Task struct {
	ID             string
	CronExpression string
	Message        json.RawMessage
}

func (t *Task) Unmarshall(raw []byte) error {
	return json.Unmarshal(raw, t)
}

func (t *Task) Marshall() ([]byte, error) {
	return json.Marshal(t)
}
