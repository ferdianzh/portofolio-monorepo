package user

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(router fiber.Router) {
	repo := NewRepository()
	handler := NewHandler(repo)

	router.Post("/", handler.Create)
	router.Get("/", handler.FindAll)
	router.Get("/:id", handler.FindOne)
	router.Put("/:id", handler.Update)
	router.Delete(":id", handler.Delete)
}
