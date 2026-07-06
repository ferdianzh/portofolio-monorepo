package user

import (
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/utils"
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{
		repo,
	}
}

func (h *Handler) Create(c fiber.Ctx) error {
	var dto CreateUserDTO

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

	user := dto.ToModel()
	user.Password = utils.HashPassword(user.Password)

	err := h.repo.Create(&user)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(user)
}

func (h *Handler) FindAll(c fiber.Ctx) error {
	users, err := h.repo.FindAll()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(users)
}

func (h *Handler) FindOne(c fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.repo.FindOne(id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(user)
}

func (h *Handler) Update(c fiber.Ctx) error {
	var dto UpdateUserDTO

	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	if err := utils.ValidateStruct(dto); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err,
		})
	}

	id := c.Params("id")

	user := dto.ToModel()
	
	if dto.Password != "" {
		user.Password = utils.HashPassword(user.Password)
	}
	
	err := h.repo.Update(&user, id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(user)
}

func (h *Handler) Delete(c fiber.Ctx) error {
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

