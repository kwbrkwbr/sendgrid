package handler

import (
	"net/http"
	"sendgrid/internal/echo/context"
)

func Healtz(c *context.MyContext) error {
	return c.String(http.StatusOK, "hc")
}
