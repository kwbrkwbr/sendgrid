package handler

import (
	"encoding/json"
	"net/http"
	"sendgrid/internal/echo/context"
)

type MailRequest struct {
	From   string `json:"from,omitempty" validate:"required,email"`
	To     string `json:"to,omitempty" validate:"required,email"`
	Title  string `json:"title,omitempty" validate:"required"`
	Body   string `json:"body,omitempty" validate:"required"`
	Params string `json:"params,omitempty" validate:"required"`
}

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

func Mail(c *context.MyContext) error {
	p := new(MailRequest)
	if err := c.BindValidate(p); err != nil {
		return c.ErrorValidation(err)
	}
	// パラメータの設定
	parse := make(map[string]string)
	if err := json.Unmarshal([]byte(p.Params), &parse); err != nil {
		return c.ErrorInternalServer(err)
	}
	res, err := c.MailDynamic(p.From, p.To, p.Title, p.Body, parse)
	if err != nil {
		return c.ErrorInternalServer(err)
	} else if res.StatusCode >= http.StatusBadRequest {
		return c.ErrorSendGridResponse(res.StatusCode, res.Body)
	}

	return c.SuccessAccepted(res.StatusCode, res.Body)
}

func MailPubSub(c *context.MyContext) error {
	m := new(PubSubMessage)
	if err := c.BindValidate(m); err != nil {
		return c.ErrorValidation(err)
	}
	p := new(MailRequest)
	if err := json.Unmarshal(m.Message.Data, p); err != nil {
		return c.ErrorInternalServer(err)
	}
	if err := c.Validate(p); err != nil {
		return c.ErrorValidation(err)
	}
	// パラメータの設定
	parse := make(map[string]string)
	if err := json.Unmarshal([]byte(p.Params), &parse); err != nil {
		return c.ErrorInternalServer(err)
	}
	res, err := c.MailDynamic(p.From, p.To, p.Title, p.Body, parse)
	if err != nil {
		return c.ErrorInternalServer(err)
	} else if res.StatusCode >= http.StatusBadRequest {
		return c.ErrorSendGridResponse(res.StatusCode, res.Body)
	}

	return c.SuccessAccepted(res.StatusCode, res.Body)
}
