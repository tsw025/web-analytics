package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/tsw025/web_analytics/internal/schemas"
	"github.com/tsw025/web_analytics/internal/services"
	"net/http"
)

type AuthHandler struct {
	authService  services.AuthService
	tokenService services.AuthTokenService
}

func NewAuthHandler(authService services.AuthService, tokenService services.AuthTokenService) *AuthHandler {
	return &AuthHandler{
		authService:  authService,
		tokenService: tokenService,
	}
}

func (l *AuthHandler) LogIn(c echo.Context) error {
	req := new(schemas.LoginRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	user, err := l.authService.Authenticate(req.Username, req.Password)
	if err != nil {
		return err
	}

	//Generate token
	token, err := l.tokenService.GenerateToken(user.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (l *AuthHandler) Register(c echo.Context) error {
	req := new(schemas.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	user, err := l.authService.Register(req.Username, req.Password)
	if err != nil {
		return err
	}

	//Generate token
	token, err := l.tokenService.GenerateToken(user.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
