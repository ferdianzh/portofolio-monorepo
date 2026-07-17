package user

import (
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/response"
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
		return response.Error(c, 400, "invalid request", err)
	}

	if err := utils.ValidateStruct(dto); err != nil {
		return response.Error(c, 400, "invalid request", err)
	}

	user := dto.ToModel()
	user.Password = utils.HashPassword(user.Password)

	err := h.repo.Create(&user)

	if err != nil {
		return response.Error(c, 500, "internal server error", err)
	}

	return response.Success(c, 201, "user created", user)
}

func (h *Handler) FindAll(c fiber.Ctx) error {
	users, err := h.repo.FindAll()

	if err != nil {
		return response.Error(c, 500, "internal server error", err)
	}

	return response.Success(c, 200, "users retrieved", users)
}

func (h *Handler) FindOne(c fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.repo.FindOne(id)

	if err != nil {
		return response.Error(c, 500, "internal server error", err)
	}

	return response.Success(c, 200, "user retrieved", user)
}

func (h *Handler) Update(c fiber.Ctx) error {
	var dto UpdateUserDTO

	if err := c.Bind().Body(&dto); err != nil {
		return response.Error(c, 400, "invalid request", err)
	}

	if err := utils.ValidateStruct(dto); err != nil {
		return response.Error(c, 400, "invalid request", err)
	}

	id := c.Params("id")

	user := dto.ToModel()
	
	if dto.Password != "" {
		user.Password = utils.HashPassword(user.Password)
	}
	
	err := h.repo.Update(&user, id)

	if err != nil {
		return response.Error(c, 500, "internal server error", err)
	}

	return response.Success(c, 200, "user updated", user)
}

func (h *Handler) Delete(c fiber.Ctx) error {
	id := c.Params("id")

	err := h.repo.Delete(id)

	if err != nil {
		return response.Error(c, 500, "internal server error", err)
	}
	
	return response.Success(c, 200, "user deleted", fiber.Map{
		"affected": 1,
	})
}

