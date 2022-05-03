package usecase

import (
	usecase "healthchecker/usecase/component"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(usecase.NewCheckComponentUseCase),
)
