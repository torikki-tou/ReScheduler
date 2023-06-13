package task

import (
	"errors"
	"github.com/gofrs/uuid"
	"github.com/robfig/cron"
	taskRepository "github.com/torikki-tou/ReScheduler/internal/repositories/task"
	repositoryDto "github.com/torikki-tou/ReScheduler/internal/repositories/task/dto"
	"github.com/torikki-tou/ReScheduler/internal/services/task/dto"
	"time"
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

	sch, err := cron.Parse(req.CronExpression)
	if err != nil {
		return nil, err
	}

	if err := s.repository.Create(&repositoryDto.CreateRequest{
		ID:             task.ID,
		Score:          sch.Next(time.Now()).Unix(),
		CronExpression: task.CronExpression,
		Message:        task.Message,
	}); err != nil {
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
func (s *Service) GetReady(req *dto.GetReadyRequest) (*dto.GetReadyResponse, error) {
	res, err := s.repository.SearchByScore(&repositoryDto.SearchByScoreRequest{MaxScore: req.NowTime.Unix()})
	if err != nil {
		return nil, err
	}

	if len(res.Tasks) == 0 {
		return &dto.GetReadyResponse{Tasks: []dto.Task{}}, nil
	}

	var tasks = make([]dto.Task, 0, len(res.Tasks))
	for _, task := range res.Tasks {

		if err := s.repository.Delete(&repositoryDto.DeleteRequest{
			ID: task.ID,
		}); err != nil {
			println(err.Error())
		}

		sch, err := cron.Parse(task.CronExpression)
		if err != nil {
			return nil, err
		}

		if err := s.repository.Create(&repositoryDto.CreateRequest{
			ID:             task.ID,
			Score:          sch.Next(time.Unix(task.Score, 0)).Unix(),
			CronExpression: task.CronExpression,
			Message:        task.Message,
		}); err != nil {
			println(err.Error())
		}

		tasks = append(tasks, *s.fromRepository(&task))
	}

	return &dto.GetReadyResponse{Tasks: tasks}, nil
}

func (s *Service) fromRepository(req *repositoryDto.Task) *dto.Task {
	return &dto.Task{
		ID:             req.ID,
		CronExpression: req.CronExpression,
		Message:        req.Message,
	}
}
