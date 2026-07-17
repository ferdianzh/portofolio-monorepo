package role

import (
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/permission"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(router fiber.Router) {
	repo := NewRepository()
	permissionRepo := permission.NewRepository()
	handler := NewHandler(repo, permissionRepo)

	group := router.Group("roles")

	group.Post("/", handler.Create)
	group.Get("/", handler.FindAll)
	group.Get("/:id", handler.FindOne)
	group.Put("/:id", handler.Update)
	group.Delete("/:id", handler.Delete)
	group.Get("/:id/permissions", handler.FindAllPermissions)
	group.Put("/:id/permissions", handler.ReplacePermissions)
}
