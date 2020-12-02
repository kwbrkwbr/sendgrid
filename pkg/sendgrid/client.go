package sendgrid

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"net/http"
	"sendgrid/pkg"
)

type Env struct {
	ApiKey   string `validate:"required"`
	Host     string `default:"https://api.sendgrid.com"`
	Endpoint string `default:"/v3/mail/send"`
}

type SgMailer interface {
	MailDynamic(from, to, title, body string, sub map[string]string) (*rest.Response, error)
}

type Client struct {
	Env *Env
}

type Email struct {
	V string `json:"email"`
}

type Content struct {
	T string `json:"type"`
	V string `json:"value"`
}

type Personalization struct {
	To            []Email           `json:"to"`
	Substitutions map[string]string `json:"substitutions"`
	Subject       string            `json:"subject"`
}

type Body struct {
	Personalizations []Personalization `json:"personalizations"`
	From             Email             `json:"from"`
	Content          []Content         `json:"content"`
}

func NewClient(e *pkg.Env) SgMailer {
	env := &Env{
		ApiKey:   e.SendgridApiKey,
		Host:     e.SendgridHost,
		Endpoint: e.SendgridEndpoint,
	}
	return Client{
		Env: env,
	}
}

func (s Client) MailDynamic(from, to, title, body string, sub map[string]string) (*rest.Response, error) {
	// パラメータの設定
	// 変数の埋め込みは `${xxx}` 形式固定にする
	for k, v := range sub {
		nk := fmt.Sprintf("${%s}", k)
		sub[nk] = v
	}

	// sendgridに送信
	request := sendgrid.GetRequest(s.Env.ApiKey,
		s.Env.Endpoint,
		s.Env.Host)
	request.Method = http.MethodPost
	request.Headers["Content-Type"] = "application/json"
	rb := Body{
		Personalizations: []Personalization{
			{
				To: []Email{
					{
						V: to,
					},
				},
				Substitutions: sub,
				Subject:       title,
			},
		},
		From: Email{V: from},
		Content: []Content{
			{
				T: "text/plain", //一旦これで固定する
				V: body,
			},
		},
	}
	bytes, err := json.Marshal(rb)
	if err != nil {
		return nil, err
	}
	request.Body = bytes

	// sendgridの処理結果チェック
	res, err := sendgrid.API(request)
	if err != nil {
		return nil, err
	}
	// responseをそのまま渡すことにする。
	//else if res.StatusCode >= http.StatusBadRequest {
	//	return res, nil
	//}
	return res, nil
}
