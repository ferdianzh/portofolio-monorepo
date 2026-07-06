package project

import (
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/gosimple/slug"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) CreateProject(c fiber.Ctx) error {
	var dto CreateProjectDTO

	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	if err := utils.ValidateStruct(dto); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err,
		})
	}

	project := dto.ToModel()
	project.Slug = slug.Make(project.Title)

	err := h.repo.Create(&project)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(project)
}

func (h *Handler) GetAllProjects(c fiber.Ctx) error {
	projects, err := h.repo.FindAll()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(projects)
}

func (h *Handler) GetOneProject(c fiber.Ctx) error {
	id := c.Params("id")

	project, err := h.repo.FindOne(id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(project)
}

func (h *Handler) UpdateProject(c fiber.Ctx) error {
	var dto UpdateProjectDTO

	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	if err := utils.ValidateStruct(dto); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err,
		})
	}

	id := c.Params("id")

	project := dto.ToModel()
	project.Slug = slug.Make(project.Title)

	err := h.repo.Update(&project, id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(project)
}

func (h *Handler) DeleteProject(c fiber.Ctx) error {
	id := c.Params("id")

	err := h.repo.Delete(id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err,
		}) 
	}
	
	return c.JSON(fiber.Map{
		"affected": 1,
	})
}
