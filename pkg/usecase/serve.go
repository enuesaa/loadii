package usecase

import (
	"mime"
	"path/filepath"
	"strings"

	"github.com/enuesaa/tryserve/pkg/repository"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func Serve(repos repository.Repos, basepath string) error {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/*", func(c fiber.Ctx) error {
		path := c.Path() // like `/`
		path = filepath.Join(basepath, path)

		// TODO: This behavior should be changed with flag.
		if strings.HasSuffix(path, "/") {
			path = filepath.Join(path, "index.html")
		}
		if ext := filepath.Ext(path); ext == "" {
			path = path + ".html"
		}

		data, err := repos.Fs.Read(path)
		if err != nil {
			return err
		}

		ext := filepath.Ext(path)
		mimeType := mime.TypeByExtension(ext)
		c.Set(fiber.HeaderContentType, mimeType)

		return c.SendString(string(data))
	})

	return app.Listen(":3000")
}
