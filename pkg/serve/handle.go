package serve

import (
	"fmt"
	"mime"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func (ctl *Servectl) handleMainRoute(c fiber.Ctx) error {
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
}
