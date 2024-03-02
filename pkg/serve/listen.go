package serve

import (
	"fmt"
	"mime"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func (ctl *Servectl) Listen() error {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/*", func(c fiber.Ctx) error {
		path := c.Path() // like `/`

		// TODO: This behavior should be changed with flag.
		if strings.HasSuffix(path, "/") {
			path = filepath.Join(path, "index.html")
		}
		if ext := filepath.Ext(path); ext == "" {
			path = path + ".html"
		}
		path = filepath.Join(ctl.Basepath, path)
		fmt.Printf("path: %s originalPath: %s\n", path, c.Path())

		data, err := ctl.repos.Fs.Read(path)
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}

		ext := filepath.Ext(path)
		mimeType := mime.TypeByExtension(ext)
		c.Set(fiber.HeaderContentType, mimeType)

		return c.SendString(string(data))
	})

	listenConfig := fiber.ListenConfig{
		DisableStartupMessage: true,
	}

	addr := fmt.Sprintf(":%d", ctl.Port)
	fmt.Printf("Listening on %s\n", addr)

	return app.Listen(addr, listenConfig)
}
