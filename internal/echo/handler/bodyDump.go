package handler

import (
	"github.com/labstack/echo/v4"
)

func BodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	c.Logger().Infof("request:%s", string(reqBody))
	c.Logger().Infof("response:%s", string(resBody))
}
