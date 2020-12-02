// +build wireinject

package main

import (
	"github.com/google/wire"
	"sendgrid/internal/fiber"
)

func InitDI() *fiber.Server {
	wire.Build(
		fiber.NewServer,
	)

	return &fiber.Server{}
}

func InitMockDI() *fiber.Server {
	wire.Build(
		fiber.NewServer,
	)

	return &fiber.Server{}
}
