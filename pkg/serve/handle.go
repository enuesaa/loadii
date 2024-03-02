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

	readpath := ctl.convertPath(path)
	fmt.Printf("path: %s, looking: %s\n", path, readpath)

	data, err := ctl.repos.Fs.Read(readpath)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	c.Set(fiber.HeaderContentType, ctl.judgeMimeType(path))

	return c.SendString(string(data))
}

func (ctl *Servectl) convertPath(path string) string {
	// TODO: This behavior should be changed with flag.
	if strings.HasSuffix(path, "/") {
		path = filepath.Join(path, "index.html")
	}
	if ext := filepath.Ext(path); ext == "" {
		path = path + ".html"
	}
	return filepath.Join(ctl.Basepath, path)
}

func (ctl *Servectl) judgeMimeType(path string) string {
	ext := filepath.Ext(path)
	return mime.TypeByExtension(ext)
}
