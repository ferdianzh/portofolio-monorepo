package role

import (
	"strings"

	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/permission"
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/response"
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
		return response.Error(c, 400, "invalid request", err)
	}

	if err := utils.ValidateStruct(dto); err != nil {
		return response.Error(c, 400, "invalid request", err)
	}

	role := dto.ToModel()

	err := h.repo.Create(&role)

	if err != nil {
		return response.Error(c, 500, "internal server error", err)
	}

	return response.Success(c, 201, "role created", role)
}

func (h *Handler) FindAll(c fiber.Ctx) error {
	roles, err := h.repo.FindAll()

	if err != nil {
		return response.Error(c, 500, "internal server error", err)
	}

	return response.Success(c, 200, "roles retrieved", roles)
}

func (h *Handler) FindOne(c fiber.Ctx) error {
	id := c.Params("id")

	role, err := h.repo.FindOne(id)

	if err != nil {
		return response.Error(c, 500, "internal server error", err)
	}

	return response.Success(c, 200, "role retrieved", role)
}

func (h *Handler) Update(c fiber.Ctx) error {
	var dto UpdateRoleDTO

	if err := c.Bind().Body(&dto); err != nil {
		return response.Error(c, 400, "invalid request", err)
	}

	if err := utils.ValidateStruct(dto); err != nil {
		return response.Error(c, 400, "invalid request", err)
	}

	id := c.Params("id")

	role := dto.ToModel()
	
	err := h.repo.Update(&role, id)

	if err != nil {
		return response.Error(c, 500, "internal server error", err)
	}

	return response.Success(c, 200, "role updated", role)
}

func (h *Handler) Delete(c fiber.Ctx) error {
	id := c.Params("id")

	err := h.repo.Delete(id)

	if err != nil {
		return response.Error(c, 500, "internal server error", err)
	}
	
	return response.Success(c, 200, "role deleted", fiber.Map{
		"affected": 1,
	})
}

func (h *Handler) FindAllPermissions(c fiber.Ctx) error {
	id := c.Params("id")

	permissions, err := h.permissionRepo.FindAll(permission.FindAllOptions{
		RoleID: id,
	})

	if err != nil {
		return response.Error(c, 500, "internal server error", err)
	}

	grouped := c.Query("grouped")

	if grouped == "1" {
		var groupedPermissions = make(map[string][]permission.Permission)
		for _, p := range permissions {
			key, _, _ := strings.Cut(p.Alias, ".")
			groupedPermissions[key] = append(groupedPermissions[key], p)
		}
		return response.Success(c, 200, "permissions retrieved", groupedPermissions)
	}

	return response.Success(c, 200, "permissions retrieved", permissions)
}

func (h *Handler) ReplacePermissions(c fiber.Ctx) error {
	var dto ReplacePermissionsDTO
	id := c.Params("id")

	if err := c.Bind().Body(&dto); err != nil {
		return response.Error(c, 400, "invalid request", err)
	}

	if err := utils.ValidateStruct(dto); err != nil {
		return response.Error(c, 400, "invalid request", err)
	}

	err := h.repo.ReplacePermissions(dto.PermissionIds, id)
	if err != nil {
		return response.Error(c, 500, "internal server error", err)
	}

	return h.FindAllPermissions(c)
}
