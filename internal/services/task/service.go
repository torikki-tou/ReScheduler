package task

import (
	"errors"
	"github.com/gofrs/uuid"
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

func (s *Service) Get(req *dto.GetRequest) (*dto.GetResponse, error) {

	res, err := s.repository.Get(&repositoryDto.GetRequest{ID: req.ID})
	if err != nil {
		return nil, err
	}

	return &dto.GetResponse{Task: *s.fromRepository(&res.Task)}, nil
}

func (s *Service) Create(req *dto.CreateRequest) (*dto.CreateResponse, error) {

	var id string
	if req.ID != nil {
		id = *req.ID

		task, err := s.repository.Get(&repositoryDto.GetRequest{ID: id})
		if err != nil {
			return nil, err
		}

		if task != nil {
			return nil, errors.New("exists")
		}
	} else {
		uuidV6, err := uuid.NewV6()
		if err != nil {
			return nil, err
		}

		id = uuidV6.String()
	}

	task := dto.Task{
		ID:             id,
		CronExpression: req.CronExpression,
		Message:        req.Message,
	}

	err := s.repository.Create(&repositoryDto.CreateRequest{
		ID:             task.ID,
		CronExpression: task.CronExpression,
		Message:        task.Message,
	})
	if err != nil {
		return nil, err
	}

	return &dto.CreateResponse{Task: task}, nil
}

func (s *Service) Update(req *dto.UpdateRequest) (*dto.UpdateResponse, error) {
	return &dto.UpdateResponse{}, nil
}

func (s *Service) Delete(req *dto.DeleteRequest) (*dto.DeleteResponse, error) {

	task, err := s.Get(&dto.GetRequest{ID: req.ID})
	if err != nil {
		return nil, err
	}

	if err := s.repository.Delete(&repositoryDto.DeleteRequest{ID: req.ID}); err != nil {
		return nil, err
	}

	return &dto.DeleteResponse{Task: task.Task}, nil
}

func (s *Service) fromRepository(req *repositoryDto.Task) *dto.Task {
	return &dto.Task{
		ID:             req.ID,
		CronExpression: req.CronExpression,
		Message:        req.Message,
	}
}
