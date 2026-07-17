package user

import (
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/middleware"
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/permission"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(router fiber.Router) {
	repo := NewRepository()
	handler := NewHandler(repo)

	group := router.Group("users")

	group.Post("/", middleware.JWTProtected(), middleware.RequirePermission(string(permission.UserCreate)), handler.Create)
	group.Get("/", middleware.JWTProtected(), middleware.RequirePermission(string(permission.UserRead)), handler.FindAll)
	group.Get("/:id", middleware.JWTProtected(), middleware.RequirePermission(string(permission.UserRead)), handler.FindOne)
	group.Put("/:id", middleware.JWTProtected(), middleware.RequirePermission(string(permission.UserUpdate)), handler.Update)
	group.Delete("/:id", middleware.JWTProtected(), middleware.RequirePermission(string(permission.UserDelete)), handler.Delete)
}
