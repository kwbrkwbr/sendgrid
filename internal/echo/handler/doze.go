package handler

import (
	"net/http"
	"sendgrid/internal/echo/context"
	"time"
)

type DozeRequest struct {
	Sec int `json:"sec"`
}

func Doze(c *context.LspmContext) error {
	c.Logger().Debug("doze...")
	r := new(DozeRequest)
	c.BindValidate(r)
	time.Sleep(time.Duration(r.Sec) * time.Second)
	return c.String(http.StatusOK, "wake up")
}
