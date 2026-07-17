package role

import (
	"strings"

	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/permission"
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/utils"
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	repo *Repository
	permissionRepo *permission.Repository
}

func NewHandler(repo *Repository, permissionRepo *permission.Repository) *Handler {
	return &Handler{
		repo,
		permissionRepo,
	}
}

func (h *Handler) Create(c fiber.Ctx) error {
	var dto CreateRoleDTO

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

	role := dto.ToModel()

	err := h.repo.Create(&role)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(role)
}

func (h *Handler) FindAll(c fiber.Ctx) error {
	roles, err := h.repo.FindAll()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(roles)
}

func (h *Handler) FindOne(c fiber.Ctx) error {
	id := c.Params("id")

	role, err := h.repo.FindOne(id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(role)
}

func (h *Handler) Update(c fiber.Ctx) error {
	var dto UpdateRoleDTO

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

	role := dto.ToModel()
	
	err := h.repo.Update(&role, id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(role)
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

func (h *Handler) FindAllPermissions(c fiber.Ctx) error {
	id := c.Params("id")

	permissions, err := h.permissionRepo.FindAll(permission.FindAllOptions{
		RoleID: id,
	})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	grouped := c.Query("grouped")

	if grouped == "1" {
		var groupedPermissions = make(map[string][]permission.Permission)
		for _, p := range permissions {
			key, _, _ := strings.Cut(p.Alias, ".")
			groupedPermissions[key] = append(groupedPermissions[key], p)
		}
		return c.JSON(groupedPermissions)
	}

	return c.JSON(permissions)
}

func (h *Handler) ReplacePermissions(c fiber.Ctx) error {
	var dto ReplacePermissionsDTO
	id := c.Params("id")

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

	err := h.repo.ReplacePermissions(dto.PermissionIds, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return h.FindAllPermissions(c)
}
