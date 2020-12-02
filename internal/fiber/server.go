package fiber

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App *fiber.App
}

func NewServer() *Server {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("hc")
	})
	return &Server{
		App: app,
	}
}

func (s Server) Run() {
	s.App.Listen(":3000")
}
