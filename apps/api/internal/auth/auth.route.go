package auth

import (
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/middleware"
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/user"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(router fiber.Router) {
	userRepo := user.NewRepository()
	handler := NewHandler(userRepo)

	group := router.Group("auth")

	group.Post("/login", handler.Login)
	group.Get("/user", middleware.JWTProtected(), handler.findUserDetail)
}
