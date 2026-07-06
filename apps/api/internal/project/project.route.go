package project

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(router fiber.Router) {
	repo := NewRepository()
	handler := NewHandler(repo)

	router.Post("/", handler.CreateProject)
	router.Get("/", handler.GetAllProjects)
	router.Get("/:id", handler.GetOneProject)
	router.Put("/:id", handler.UpdateProject)
	router.Delete("/:id", handler.DeleteProject)
}
