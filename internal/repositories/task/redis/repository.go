package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/torikki-tou/ReScheduler/internal/repositories/task/dto"
	"strconv"
)

type Repository struct {
	cli *redis.Client
}

func New(config Config) *Repository {
	addr := config.RedisHost() + `:` + strconv.Itoa(config.RedisPort())

	return &Repository{
		cli: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}
}

func (r *Repository) Get(req *dto.GetRequest) (*dto.GetResponse, error) {
	res, err := r.cli.Get(context.Background(), req.ID).Bytes()
	if err != nil {
		return nil, err
	}

	var task dto.Task
	if err := task.Unmarshall(res); err != nil {
		return nil, err
	}

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

	marshalledTask, err := task.Marshall()
	if err != nil {
		return err
	}

	return r.cli.Set(context.Background(), task.ID, marshalledTask, 0).Err()
}

func (r *Repository) Update(req *dto.UpdateRequest) error {
	return nil
}

func (r *Repository) Delete(req *dto.DeleteRequest) error {
	_ = r.cli.Del(context.Background(), req.ID)

	return nil
}
