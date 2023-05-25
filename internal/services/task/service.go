package task

import "github.com/torikki-tou/ReScheduler/internal/services/task/dto"

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) Get(req *dto.GetRequest) *dto.GetResponse {
	return &dto.GetResponse{}
}

func (s *Service) Search(req *dto.SearchRequest) *dto.SearchResponse {
	return &dto.SearchResponse{}
}

func (s *Service) Create(req *dto.CreateRequest) *dto.CreateResponse {
	return &dto.CreateResponse{}
}

func (s *Service) Update(req *dto.UpdateRequest) *dto.UpdateResponse {
	return &dto.UpdateResponse{}
}

func (s *Service) Delete(req *dto.DeleteRequest) *dto.DeleteResponse {
	return &dto.DeleteResponse{}
}
