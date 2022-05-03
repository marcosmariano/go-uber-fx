package ports

import (
	"healthchecker/ports/config"

	"go.uber.org/fx"
)

var Module = fx.Options(
	config.Module,
)
