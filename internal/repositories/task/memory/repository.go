package memory

import (
	"github.com/torikki-tou/ReScheduler/internal/repositories/task/dto"
	"github.com/wangjia184/sortedset"
)

type Repository struct {
	set *sortedset.SortedSet
}

func New() *Repository {
	return &Repository{
		set: sortedset.New(),
	}
}

func (r *Repository) Get(req *dto.GetRequest) (*dto.GetResponse, error) {
	node := r.set.GetByKey(req.ID)
	if node == nil {
		return nil, nil
	}

	task := node.Value.(dto.Task)

	return &dto.GetResponse{Task: dto.Task{
		ID:             task.ID,
		CronExpression: task.CronExpression,
		Message:        task.Message,
	}}, nil
}

func (r *Repository) Create(req *dto.CreateRequest) error {
	task := dto.Task{
		ID:             req.ID,
		CronExpression: req.CronExpression,
		Message:        req.Message,
	}
	r.set.AddOrUpdate(task.ID, sortedset.SCORE(req.Score), task)

	return nil
}

func (r *Repository) Update(req *dto.UpdateRequest) error {
	return nil
}

func (r *Repository) Delete(req *dto.DeleteRequest) error {
	_ = r.set.Remove(req.ID)
	return nil
}
