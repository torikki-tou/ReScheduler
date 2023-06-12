package task

import "github.com/torikki-tou/ReScheduler/internal/repositories/task/dto"

type Repository interface {
	Get(req *dto.GetRequest) (*dto.GetResponse, error)
	Create(req *dto.CreateRequest) error
	Update(req *dto.UpdateRequest) error
	Delete(req *dto.DeleteRequest) error
}
