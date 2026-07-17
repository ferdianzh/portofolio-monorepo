package middleware

import (
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/permission"
	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func RequirePermission(permissionAlias string) fiber.Handler {
	return func(c fiber.Ctx) error {
		u := jwtware.FromContext(c)
		if u == nil {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		claims := u.Claims.(jwt.MapClaims)
		roleId, ok := claims["role"].(string)
		if !ok || roleId == "" {
			return c.Status(403).JSON(fiber.Map{
				"message": "Forbidden",
			})
		}

		repo := permission.NewRepository()
		permissions, err := repo.FindAll(permission.FindAllOptions{
			RoleID: roleId,
		})

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		}

		hasPermission := false
		for _, p := range permissions {
			if p.Alias == permissionAlias {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			return c.Status(403).JSON(fiber.Map{
				"message": "Forbidden",
			})
		}

		return c.Next()
	}
}
