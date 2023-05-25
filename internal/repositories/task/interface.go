package task

import "github.com/torikki-tou/ReScheduler/internal/repositories/task/dto"

type Repository interface {
	Get(req *dto.GetRequest) *dto.GetResponse
	Search(req *dto.SearchRequest) *dto.SearchResponse
	Create(req *dto.CreateRequest) *dto.CreateResponse
	Update(req *dto.UpdateRequest) *dto.UpdateResponse
	Delete(req *dto.DeleteRequest) *dto.DeleteResponse
}
