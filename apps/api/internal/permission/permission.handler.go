package permission

import (
	"strings"

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

func (h *Handler) FindAll(c fiber.Ctx) error {
	permissions, err := h.repo.FindAll()

	grouped := c.Query("grouped")

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if grouped == "1" {
		var groupedPermissions = make(map[string][]Permission)
		for _, permission := range permissions {
			key, _, _ := strings.Cut(permission.Alias, ".")
			groupedPermissions[key] = append(groupedPermissions[key], permission)
		}
		return c.JSON(groupedPermissions)
	}

	return c.JSON(permissions)
}
