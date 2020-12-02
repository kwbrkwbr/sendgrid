// +build mock

package main

import (
	s "sendgrid/internal/fiber"
)

func NewServer() *s.Server {
	return InitMockDI()
}
