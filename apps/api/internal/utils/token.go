package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(sub string, roleId *string) (string, error) {
	claims := jwt.MapClaims{
		"sub": sub,
		"iat": time.Now().Add(time.Hour * 1).Unix(),
	}
	
	if roleId != nil {
		claims["role"] = *roleId
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "error", err
	}

	return t, nil
}

func VerifyToken(token string) (bool, error) {
	jwt, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return false, err
	}

	return jwt.Valid, nil
}
