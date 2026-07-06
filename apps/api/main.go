package main

import (
	"log"

	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/database"
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/migrations"
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/project"
	"github.com/gofiber/fiber/v3"
)

func main() {
    app := fiber.New()

    database.Connect()

    migrations.Migrate()

    api := app.Group("/api")

    project.RegisterRoutes(
        api.Group("/projects"),
    )

    app.Get("/", func(c fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    log.Fatal(app.Listen(":3000"))
}
