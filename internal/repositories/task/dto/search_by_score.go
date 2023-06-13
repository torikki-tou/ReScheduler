package dto

type (
	SearchByScoreRequest struct {
		MaxScore int64
	}

	SearchByScoreResponse struct {
		Tasks []Task
	}
)
