package echo

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"gopkg.in/tylerb/graceful.v1"
	"os"
	myContext "sendgrid/internal/echo/context"
	"sendgrid/internal/echo/handler"
	"sendgrid/pkg"
	myLog "sendgrid/pkg/log"
	"sendgrid/pkg/sendgrid"
	"time"
)

type Server struct {
	a *echo.Echo
	e *pkg.Env
	c sendgrid.SgMailer
}

func NewServer(e *pkg.Env, c sendgrid.SgMailer) *Server {
	a := echo.New()
	a.Server.Addr = fmt.Sprintf(":%d", e.Port)
	return &Server{
		a: a,
		e: e,
		c: c,
	}
}

func (s Server) Run() {
	switch s.e.Env {
	case "test":
		s.a.Logger.SetLevel(log.DEBUG)
	default:
		s.a.HideBanner = true
		//s.a.HidePort = true // これをなくすと起動したかわかりにくいのでつけておく
		s.a.Logger.SetLevel(log.INFO)
	}

	s.a.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10,
		LogLevel:  0,
	}))
	s.a.Use(middleware.RequestID())
	s.a.Use(middleware.Logger())
	s.a.Validator = &myContext.Validator{Validator: validator.New()}

	// ログにリクエストIDを仕込むもの。共通でログに出したいならここを改造
	s.a.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			l := logrus.New()
			l.SetFormatter(&logrus.JSONFormatter{})
			l.SetOutput(os.Stdout)
			myLogger := &myLog.MyLogger{
				Logger:    l,
				RequestID: id,
			}
			c.SetLogger(myLogger)
			return h(c)
		}
	})
	// request/responseのbodyのロギング
	s.a.Use(middleware.BodyDump(handler.BodyDumpHandler))

	// LspmContextをラップする
	s.a.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return h(&myContext.MyContext{Context: c, Env: s.e, SgMailer: s.c})
		}
	})

	// health check
	s.a.GET("/healthz", myContext.Wrap(handler.Healtz))
	// health check
	s.a.GET("/doze", myContext.Wrap(handler.Doze))
	// 一般的なメール送信の利用を想定
	s.a.POST("/mail", myContext.Wrap(handler.Mail))
	// 一般的なメール送信の利用を想定(パラメータ埋め込みあり)
	s.a.POST("/mail/pubsub", myContext.Wrap(handler.MailPubSub))

	// graceful shutdown
	//if err := graceful.ListenAndServe(s.a.Server, 10*time.Second); err != nil {
	//	s.a.Logger.Error(err)
	//}
	if err := graceful.ListenAndServe(s.a.Server, 10*time.Second); err != nil {
		s.a.Logger.Error(err)
	}
}
