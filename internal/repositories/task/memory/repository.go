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

func (r *Repository) Get(req *dto.GetRequest) *dto.GetResponse {
	node := r.set.GetByKey(req.ID)
	task := node.Value.(dto.Task)

	return &dto.GetResponse{Task: dto.Task{
		ID:             task.ID,
		CronExpression: task.CronExpression,
		Message:        task.Message,
	}}
}

func (r *Repository) Search(req *dto.SearchRequest) *dto.SearchResponse {
	var limit = -1
	if req.Limit != nil {
		limit = *req.Limit
	}

	nodes := r.set.GetByRankRange(1, limit, false)

	var tasks []dto.Task
	for _, node := range nodes {
		task := node.Value.(dto.Task)
		tasks = append(tasks, task)
	}

	return &dto.SearchResponse{Tasks: tasks}
}

func (r *Repository) Create(req *dto.CreateRequest) *dto.CreateResponse {
	task := dto.Task{
		ID:             req.ID,
		CronExpression: req.CronExpression,
		Message:        req.Message,
	}
	r.set.AddOrUpdate(task.ID, sortedset.SCORE(req.Score), task)

	return &dto.CreateResponse{Task: task}
}

func (r *Repository) Update(req *dto.UpdateRequest) *dto.UpdateResponse {
	return &dto.UpdateResponse{}
}

func (r *Repository) Delete(req *dto.DeleteRequest) *dto.DeleteResponse {
	node := r.set.Remove(req.ID)
	task := node.Value.(dto.Task)

	return &dto.DeleteResponse{Task: task}
}
