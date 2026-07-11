package serve

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func (ctl *Servectl) Addr() string {
	return fmt.Sprintf(":%d", ctl.Port)
}

func (ctl *Servectl) App() *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Get("/*", ctl.handleMainRoute)

	return app
}

func (ctl *Servectl) Listen() error {
	listenConfig := fiber.ListenConfig{
		DisableStartupMessage: true,
	}
	return ctl.App().Listen(ctl.Addr(), listenConfig)
}
