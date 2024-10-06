package auth

import (
	"github.com/labstack/echo/v4"
)

func (h *AuthHandler) RegisterRoutes(e *echo.Group) {
	authGroup := e.Group("/auth")
	authGroup.POST("/login", h.LogIn)
	authGroup.POST("/register", h.Register)
}
