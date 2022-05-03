package adapters

import (
	"healthchecker/adapters/logger"
	"healthchecker/adapters/repository"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(logger.GetLogger),
	repository.Module,
)
