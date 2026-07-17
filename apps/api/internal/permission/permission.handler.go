package permission

import (
	"strings"

	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/response"
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
		return response.Error(c, 500, "internal server error", err)
	}

	if grouped == "1" {
		var groupedPermissions = make(map[string][]Permission)
		for _, permission := range permissions {
			key, _, _ := strings.Cut(permission.Alias, ".")
			groupedPermissions[key] = append(groupedPermissions[key], permission)
		}
		return response.Success(c, 200, "permissions retrieved", groupedPermissions)
	}

	return response.Success(c, 200, "permissions retrieved", permissions)
}
