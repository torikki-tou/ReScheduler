package dto

import (
	"net/http"
	"strconv"
)

type (
	SearchRequest struct {
		Limit *int
	}

	SearchResponse struct {
		Tasks []Task `json:"tasks"`
	}
)

func (r *SearchRequest) Bind(req *http.Request) error {
	queryLimit := req.URL.Query().Get("limit")
	if queryLimit == "" {
		return nil
	}

	limit, err := strconv.Atoi(queryLimit)
	if err != nil {
		return err
	}

	r.Limit = &limit
	return nil
}
