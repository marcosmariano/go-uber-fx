package main

import (
	"healthchecker/bootstrap"

	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.Module).Run()
}
