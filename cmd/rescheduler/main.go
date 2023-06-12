package main

import (
	"github.com/torikki-tou/ReScheduler/internal/api"
	v1Api "github.com/torikki-tou/ReScheduler/internal/api/v1"
	taskV1Api "github.com/torikki-tou/ReScheduler/internal/api/v1/api/task"
	configComponent "github.com/torikki-tou/ReScheduler/internal/components/config"
	serverComponent "github.com/torikki-tou/ReScheduler/internal/components/server"
	taskRepository "github.com/torikki-tou/ReScheduler/internal/repositories/task"
	taskRepositoryMemory "github.com/torikki-tou/ReScheduler/internal/repositories/task/memory"
	taskRepositoryRedis "github.com/torikki-tou/ReScheduler/internal/repositories/task/redis"
	taskService "github.com/torikki-tou/ReScheduler/internal/services/task"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			func() *configComponent.Config {
				config := configComponent.New()
				config.SetDefaults()
				return config
			},
			func(config *configComponent.Config) taskRepository.Repository {
				switch config.GetStorageType() {
				case configComponent.StorageTypeRedis:
					return taskRepositoryRedis.New(config)
				case configComponent.StorageTypeMemory:
					return taskRepositoryMemory.New()
				default:
					return nil
				}
			},
			func(repository taskRepository.Repository) *taskService.Service {
				return taskService.New(repository)
			},
			func(service *taskService.Service) *taskV1Api.API {
				return taskV1Api.New(service)
			},
			func(task *taskV1Api.API) *v1Api.API {
				return v1Api.New(task)
			},
			func(v1 *v1Api.API) *api.API {
				return api.New(v1)
			},
			func(api *api.API, config *configComponent.Config) *serverComponent.Server {
				return serverComponent.New(api.Router(), config)
			},
		),
		fx.Invoke(
			func(server *serverComponent.Server) {
				_ = server.Run()
			},
		),
	)
}
