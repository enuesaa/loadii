package usecase

import (
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func Serve() error {
	app := fiber.New()

	app.Get("/*", func(c fiber.Ctx) error {
		path := c.Path() // like `/`
		if strings.HasSuffix(path, "/") {
			path = filepath.Join(path, "index.html")
		}
		if !strings.Contains(path, ".") {
			path = path + ".html"
		}

		f, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		fileExt := filepath.Ext(path)
		mimeType := mime.TypeByExtension(fileExt)
		c.Set(fiber.HeaderContentType, mimeType)

		return c.SendString(string(f))
	})

	return app.Listen(":3000")
}
