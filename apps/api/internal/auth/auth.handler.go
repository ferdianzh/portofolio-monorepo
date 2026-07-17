package auth

import (
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/user"
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/utils"
	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

type UserRepository interface {
    FindByEmail(email string) (*user.User, error)
		FindOne(id string) (*user.User, error)
}

type Handler struct {
	userRepo UserRepository
}

func NewHandler(userRepo *user.Repository) *Handler {
	return &Handler{
		userRepo,
	}
}

func (h *Handler) Login(c fiber.Ctx) error {
	var dto LoginDTO

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

	credErrMsg := fiber.Map{
		"message": "Invalid credentials",
	}

	user, err := h.userRepo.FindByEmail(dto.Email)
	if err != nil {
		return c.Status(401).JSON(credErrMsg)
	}

	isMatch := utils.ComparePassword(user.Password, dto.Password)
	if isMatch == false {
		return c.Status(401).JSON(credErrMsg)
	}

	var roleId *string
	if user.RoleID != nil {
		rid := user.RoleID.String()
		roleId = &rid
	}

	token, err := utils.GenerateToken(user.ID.String(), roleId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func (h *Handler) findUserDetail(c fiber.Ctx) error {
	u := jwtware.FromContext(c)
	claims := u.Claims.(jwt.MapClaims)
	
	id := claims["sub"].(string)

	user, err := h.userRepo.FindOne(id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.JSON(user)
}
