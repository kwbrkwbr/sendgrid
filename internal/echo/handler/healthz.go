package handler

import (
	"net/http"
	"sendgrid/internal/echo/context"
)

func Healtz(c *context.LspmContext) error {
	return c.String(http.StatusOK, "hc")
}
