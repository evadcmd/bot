package main

import (
	"log/slog"
	"os"

	"github.com/evadcmd/bot/pkg/api"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

//	@title			BOT
//	@version		0.0.1
//	@description	an experimental chat API
//	@contact.name	evadcmd
//	@contact.email	evadcmd@gmail.com
//	@host			localhost:5252
//	@basePath		/
//	@schemes		http
func main() {
	app := fiber.New()
	app.Use(pprof.New())
	app.Use(cors.New())
	app.Use(swagger.New(swagger.Config{FilePath: "./docs/swagger.json", BasePath: "/"}))

	app.Get("/healthz", api.ReadinessProbe)
	app.Post("/api/v0/mrkl", api.MRKL)

	if err := app.Listen(":5252"); err != nil {
		panic(err)
	}
}

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}
