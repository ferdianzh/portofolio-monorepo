package permission

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(app fiber.Router) {
	repo := NewRepository()
	handler := NewHandler(repo)

	permissionRoutes := app.Group("/permissions")

	permissionRoutes.Get("/", handler.FindAll)
}