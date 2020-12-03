package context

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"sendgrid/pkg"
	"sendgrid/pkg/sendgrid"
	"strconv"
)

type Validator struct {
	Validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}

type MyContext struct {
	echo.Context
	*pkg.Env
	sendgrid.SgMailer
}

func (c *MyContext) BindValidate(i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}
	if err := c.Validate(i); err != nil {
		return err
	}
	return nil
}

type LspmResponse struct {
	ID           string `json:"id"`              // httpStatus
	Code         int    `json:"code"`            // sendgridとしてのcodeを入れる
	ProviderCode string `json:"pcode,omitempty"` // sendgridから帰ってきたcodeを入れる
	Msg          string `json:"msg"`             // メッセージはここ
	Error        string `json:"error,omitempty"` // error.Error()をここにいれる予定
}

func (c *MyContext) RequestID() string {
	req := c.Request()
	res := c.Response()
	id := req.Header.Get(echo.HeaderXRequestID)
	if id == "" {
		id = res.Header().Get(echo.HeaderXRequestID)
	}
	return id
}

func (c *MyContext) ErrorInternalServer(err error) error {
	c.Logger().Error(err)
	r := LspmResponse{
		ID:    c.RequestID(),
		Code:  http.StatusInternalServerError,
		Msg:   "internal server error",
		Error: err.Error(),
	}
	return c.JSON(http.StatusAccepted, r)
}

func (c *MyContext) ErrorValidation(err error) error {
	c.Logger().Error(err)
	res := &LspmResponse{
		ID:    c.RequestID(),
		Code:  http.StatusBadRequest,
		Msg:   "validation error",
		Error: err.Error(),
	}
	return c.JSON(http.StatusAccepted, res)
	//return c.JSON(http.StatusAccepted, res)
}

func (c *MyContext) ErrorSendGrid(err error) error {
	c.Logger().Error(err)
	r := &LspmResponse{
		ID:           c.RequestID(),
		Code:         http.StatusInternalServerError,
		ProviderCode: "XXX", // sendgridじゃないとわからないエラーが起きた場合を想定
		Msg:          "SendGrid API error",
		Error:        err.Error(),
	}
	return c.JSON(http.StatusAccepted, r)
}

func (c *MyContext) ErrorSendGridResponse(code int, msg string) error {
	c.Logger().Errorf("sendgrid error:code=%d,msg=%s", code, msg)
	r := &LspmResponse{
		ID:           c.RequestID(),
		Code:         http.StatusInternalServerError,
		ProviderCode: strconv.Itoa(code),
		Msg:          "SendGrid Return errors",
		Error:        msg,
	}
	return c.JSON(http.StatusAccepted, r)
}

func (c *MyContext) SuccessAccepted(code int, msg string) error {
	r := &LspmResponse{
		ID:           c.RequestID(),
		Code:         http.StatusAccepted,
		ProviderCode: strconv.Itoa(code),
		Msg:          msg,
	}
	return c.JSON(http.StatusAccepted, r)
}

type handlerFunc func(c *MyContext) error

func Wrap(h handlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c.(*MyContext))
	}
}
