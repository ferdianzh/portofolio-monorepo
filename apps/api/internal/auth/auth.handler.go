package auth

import (
	"github.com/ferdianzh/portofolio-monorepo/apps/api/internal/response"
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
		return response.Error(c, 400, "invalid request", err)
	}

	if err := utils.ValidateStruct(dto); err != nil {
		return response.Error(c, 400, "invalid request", err)
	}

	user, err := h.userRepo.FindByEmail(dto.Email)
	if err != nil {
		return response.Error(c, 401, "invalid credentials", err)
	}

	isMatch := utils.ComparePassword(user.Password, dto.Password)
	if isMatch == false {
		return response.Error(c, 401, "invalid credentials", nil)
	}

	var roleId *string
	if user.RoleID != nil {
		rid := user.RoleID.String()
		roleId = &rid
	}

	token, err := utils.GenerateToken(user.ID.String(), roleId)
	if err != nil {
		return response.Error(c, 500, "internal server error", err)
	}

	return response.Success(c, 200, "login success", fiber.Map{
		"token": token,
	})
}

func (h *Handler) findUserDetail(c fiber.Ctx) error {
	u := jwtware.FromContext(c)
	claims := u.Claims.(jwt.MapClaims)
	
	id := claims["sub"].(string)

	user, err := h.userRepo.FindOne(id)

	if err != nil {
		return response.Error(c, 500, "internal server error", err)
	}

	return response.Success(c, 200, "user detail retrieved", user)
}
