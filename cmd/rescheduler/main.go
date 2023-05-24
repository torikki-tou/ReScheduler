package main

import (
	"github.com/torikki-tou/ReScheduler/internal/api"
	v1Api "github.com/torikki-tou/ReScheduler/internal/api/v1"
	taskV1Api "github.com/torikki-tou/ReScheduler/internal/api/v1/api/task"
	configComponent "github.com/torikki-tou/ReScheduler/internal/components/config"
	serverComponent "github.com/torikki-tou/ReScheduler/internal/components/server"
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
			func() *taskService.Service {
				return taskService.New()
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
