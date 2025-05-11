package v1

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	jwtSecret string
}

func NewAuthHandler(jwtSecret string) *AuthHandler {
	return &AuthHandler{jwtSecret: jwtSecret}
}

func (h *AuthHandler) Login(c echo.Context) error {
	// Пример: всегда user_id=1, без проверки пароля
	claims := jwt.MapClaims{
		"user_id": 1,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(h.jwtSecret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not generate token"})
	}
	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}
