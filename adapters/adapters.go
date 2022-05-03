package adapters

import (
	"healthchecker/adapters/logger"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(logger.GetLogger),
)
