package pkg

import (
	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	SendgridApiKey   string `envconfig:"SENDGRID_API_KEY"`
	SendgridHost     string `envconfig:"SENDGRID_HOST" default:"https://api.sendgrid.com"` // 実装当時のhost、必要に応じて変える
	SendgridEndpoint string `envconfig:"SENDGRID_ENDPOINT" default:"/v3/mail/send"`        // 実装当時のendpoint、必要に応じて変える
	Env              string `envconfig:"ENV"`
	Port             int    `envconfig:"PORT" default:"8080"` //cloud runにある環境変数
}

func NewEnv() *Env {
	e := new(Env)
	if err := envconfig.Process("", e); err != nil {
		return e
	}
	return e
}
