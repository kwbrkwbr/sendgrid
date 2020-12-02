package sendgrid

import (
	"github.com/sendgrid/rest"
	"net/http"
	"sendgrid/pkg"
)

type MockClient struct {
	Env *Env
}

type MockResponse struct {
	Code int
	Body string
}

func NewMockClient(e *pkg.Env) SgMailer {
	env := &Env{
		ApiKey:   e.SendgridApiKey,
		Host:     e.SendgridHost,
		Endpoint: e.SendgridEndpoint,
	}
	return MockClient{
		Env: env,
	}
}

func (s MockClient) MailDynamic(from, to, title, body string, sub map[string]string) (*rest.Response, error) {
	return &rest.Response{
		StatusCode: http.StatusOK,
		Body:       "mock MailDynamic",
	}, nil
}
