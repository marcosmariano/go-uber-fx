package bootstrap

import (
	"context"
	"healthchecker/adapters"
	"healthchecker/adapters/logger"
	"healthchecker/ports"
	"healthchecker/ports/config"
	"healthchecker/usecase"
	checkcomponentusecase "healthchecker/usecase/component"
	"time"

	"github.com/go-co-op/gocron"
	"go.uber.org/fx"
)

var Module = fx.Options(
	adapters.Module,
	ports.Module,
	usecase.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	logger logger.Logger,
	config config.Config,
	checkcomponent checkcomponentusecase.CheckComponentUseCase,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			go func() {
				logger.Info("Starting Healthchecker Job")
				s := gocron.NewScheduler(time.UTC)
				_, err := s.Every(config.Viper.GetInt("app.healthchecker_job.frequency_millisecond")).Millisecond().Do(
					func() {
						checkcomponent.Check()
					})
				if err != nil {
					logger.Errorf("Error on start new job")
					return
				}
				s.StartAsync()
			}()
			return nil
		},
	})
}
