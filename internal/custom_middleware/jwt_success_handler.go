// internal/middleware/jwt_success_handler.go
package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/tsw025/web_analytics/internal/echologrus"
	"github.com/tsw025/web_analytics/internal/handlers"
	"github.com/tsw025/web_analytics/internal/repositories"
	"net/http"
	"strconv"
)

// It fetches the user from the database and sets it in the Echo context.
func JWTSuccessHandler(userRepo repositories.UserRepository) func(c echo.Context) {
	return func(c echo.Context) {
		// Retrieve the JWT token from the context
		userToken, ok := c.Get("user").(*jwt.Token)
		if !ok || !userToken.Valid {
			echologrus.Logger.Info("Invalid JWT token")
			panic(handlers.NewDomainError(http.StatusUnauthorized, "Invalid JWT token", errors.New("Invalid JWT token")))
		}

		// Extract claims from the token
		sub, _ := userToken.Claims.GetSubject()
		userID, _ := strconv.ParseUint(sub, 10, 64)

		// Fetch the user from the database
		user, err := userRepo.GetByID(uint(userID))
		if err != nil {
			echologrus.Logger.Infof("User not found: %d", userID)
			panic(handlers.NewDomainError(http.StatusUnauthorized, "User not found", errors.New("User not found")))
		}

		// Log successful authentication
		echologrus.Logger.Debugf("User %s authenticated successfully", user.Username)

		// Set the user object in the Echo context
		c.Set("user", user)
	}
}
