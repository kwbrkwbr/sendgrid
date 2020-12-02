// +build mock

package main

import (
	"sendgrid/internal/echo"
)

func NewServer() *echo.Server {
	return InitMockDI()
}
