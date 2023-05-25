package dto

type (
	SearchRequest struct {
		Limit *int
	}

	SearchResponse struct {
		Tasks []Task
	}
)
