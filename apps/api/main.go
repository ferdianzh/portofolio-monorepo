package main

import (
	"log"

	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/auth"
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/database"
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/migrations"
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/project"
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/user"
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/utils"
	"github.com/gofiber/fiber/v3"
)

func main() {
    app := fiber.New()

    database.Connect()

    migrations.Migrate()

    utils.InitValidator()

    api := app.Group("/api")

    auth.RegisterRoutes(api)
    user.RegisterRoutes(api.Group("/users"))
    project.RegisterRoutes(api.Group("/projects"))

    app.Get("/", func(c fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    log.Fatal(app.Listen(":3000"))
}
