package repository

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewDynamoDbFactory),
	fx.Provide(NewComponentRepository),
)
