package worker

import (
	taskService "github.com/torikki-tou/ReScheduler/internal/services/task"
	"github.com/torikki-tou/ReScheduler/internal/services/task/dto"
	"time"
)

type Worker struct {
	tasks *taskService.Service
}

func New(tasks *taskService.Service) *Worker {
	return &Worker{
		tasks: tasks,
	}
}

func (w *Worker) Run() {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case nowTime := <-ticker.C:
			res, err := w.tasks.GetReady(&dto.GetReadyRequest{NowTime: nowTime})
			if err != nil {
				println(err.Error())
			}
			for _, task := range res.Tasks {
				println(task.ID)
			}

		}
	}
}
