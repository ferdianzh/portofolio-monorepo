package user

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(router fiber.Router) {
	repo := NewRepository()
	handler := NewHandler(repo)

	group := router.Group("users")

	group.Post("/", handler.Create)
	group.Get("/", handler.FindAll)
	group.Get("/:id", handler.FindOne)
	group.Put("/:id", handler.Update)
	group.Delete(":id", handler.Delete)
}
