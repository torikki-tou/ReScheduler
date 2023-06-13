package dto

import "time"

type (
	GetReadyRequest struct {
		NowTime time.Time
	}

	GetReadyResponse struct {
		Tasks []Task
	}
)
