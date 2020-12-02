// +build wireinject

package main

import (
	"github.com/google/wire"
	"sendgrid/internal/echo"
	"sendgrid/pkg"
	"sendgrid/pkg/sendgrid"
)

func InitDI() *echo.Server {
	wire.Build(
		pkg.NewEnv,
		sendgrid.NewClient,
		echo.NewServer,
	)
	return &echo.Server{}
}

func InitMockDI() *echo.Server {
	wire.Build(
		pkg.NewEnv,
		sendgrid.NewMockClient,
		echo.NewServer,
	)
	return &echo.Server{}
}
