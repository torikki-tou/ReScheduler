package task

import (
	taskRepository "github.com/torikki-tou/ReScheduler/internal/repositories/task"
	repositoryDto "github.com/torikki-tou/ReScheduler/internal/repositories/task/dto"
	"github.com/torikki-tou/ReScheduler/internal/services/task/dto"
)

type Service struct {
	repository taskRepository.Repository
}

func New(repository taskRepository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Get(req *dto.GetRequest) *dto.GetResponse {

	res := s.repository.Get(&repositoryDto.GetRequest{ID: req.ID})

	return &dto.GetResponse{Task: *s.fromRepository(&res.Task)}
}

func (s *Service) Search(req *dto.SearchRequest) *dto.SearchResponse {
	res := s.repository.Search(&repositoryDto.SearchRequest{Limit: req.Limit})

	var tasks = make([]dto.Task, 0, len(res.Tasks))
	for _, task := range res.Tasks {
		tasks = append(tasks, *s.fromRepository(&task))
	}

	return &dto.SearchResponse{Tasks: tasks}
}

func (s *Service) Create(req *dto.CreateRequest) *dto.CreateResponse {

	var id string
	if req.ID != nil {
		id = *req.ID
	} else {
		id = "1"
	}

	res := s.repository.Create(&repositoryDto.CreateRequest{
		ID:             id,
		CronExpression: req.CronExpression,
		Message:        req.Message,
	})

	return &dto.CreateResponse{Task: *s.fromRepository(&res.Task)}
}

func (s *Service) Update(req *dto.UpdateRequest) *dto.UpdateResponse {
	return &dto.UpdateResponse{}
}

func (s *Service) Delete(req *dto.DeleteRequest) *dto.DeleteResponse {

	res := s.repository.Delete(&repositoryDto.DeleteRequest{ID: req.ID})

	return &dto.DeleteResponse{Task: *s.fromRepository(&res.Task)}
}

func (s *Service) fromRepository(req *repositoryDto.Task) *dto.Task {
	return &dto.Task{
		ID:             req.ID,
		CronExpression: req.CronExpression,
		Message:        req.Message,
	}
}
