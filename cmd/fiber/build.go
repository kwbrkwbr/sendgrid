// +build !mock

package main

import (
	"sendgrid/internal/fiber"
)

func NewServer() *fiber.Server {
	return InitDI()
}
