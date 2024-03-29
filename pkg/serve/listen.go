package serve

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func (ctl *Servectl) Addr() string {
	return fmt.Sprintf(":%d", ctl.Port)
}

func (ctl *Servectl) Listen() error {
	app := fiber.New()

	app.Use(cors.New())
	app.Get("/*", ctl.handleMainRoute)

	listenConfig := fiber.ListenConfig{
		DisableStartupMessage: true,
	}
	return app.Listen(ctl.Addr(), listenConfig)
}
