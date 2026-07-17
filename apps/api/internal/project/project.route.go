package project

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(router fiber.Router) {
	repo := NewRepository()
	handler := NewHandler(repo)

	group := router.Group("projects")

	group.Post("/", handler.CreateProject)
	group.Get("/", handler.GetAllProjects)
	group.Get("/:id", handler.GetOneProject)
	group.Put("/:id", handler.UpdateProject)
	group.Delete("/:id", handler.DeleteProject)
}
