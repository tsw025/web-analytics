package handlers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echologrus "github.com/tsw025/web_analytics/internal/middleware"
	"github.com/tsw025/web_analytics/internal/schemas"
	"net/http"
)

// CustomError type
type CustomError struct {
	Code    int
	Message string
	Err     error
}

// Implement the error interface for CustomError
func (ce *CustomError) Error() string {
	return fmt.Sprintf("%d: %s", ce.Code, ce.Message)
}

// Unwrap allows us to use errors.Is and errors.As with CustomError
func (ce *CustomError) Unwrap() error {
	return ce.Err
}

// Helper function to create new CustomError
func NewDomainError(code int, message string, err error) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	// Handle HTTP errors like 404, 403, etc.
	if he, ok := err.(*echo.HTTPError); ok {
		var jsonErr error
		switch he.Code {
		case http.StatusNotFound:
			jsonErr = c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Resource not found",
				"error":   "The requested resource could not be found on the server.",
			})
		case http.StatusUnauthorized:
			jsonErr = c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Unauthorized access",
			})
		default:
			jsonErr = c.JSON(he.Code, map[string]interface{}{
				"message": he.Message,
			})
		}
		if jsonErr != nil {
			echologrus.Logger.Error(jsonErr)
		}
		return
	}

	// Handle validation errors
	if ve, ok := err.(validator.ValidationErrors); ok {
		var errors []schemas.ValidationErrorResponse
		for _, e := range ve {
			errors = append(errors, schemas.ValidationErrorResponse{
				Field:   e.Field(),
				Message: "failed on the " + e.Tag() + " tag",
			})
		}
		if jsonErr := c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "validation failed",
			"errors":  errors,
		}); jsonErr != nil {
			echologrus.Logger.Error(jsonErr)
		}
		return
	}

	// Handle domain errors
	if de, ok := err.(*CustomError); ok {
		if de.Code == http.StatusBadRequest {
			errors := []schemas.ValidationErrorResponse{
				{
					Field:   de.Err.Error(),
					Message: de.Message,
				},
			}
			if jsonErr := c.JSON(de.Code, map[string]interface{}{
				"message": "validation failed",
				"errors":  errors,
			}); jsonErr != nil {
				echologrus.Logger.Error(jsonErr)
			}
			return
		}
		if jsonErr := c.JSON(de.Code, map[string]interface{}{
			"message": de.Message,
			"error":   de.Err.Error(),
		}); jsonErr != nil {
			echologrus.Logger.Error(jsonErr)
		}
		return
	}

	// Handle generic internal server errors
	if jsonErr := c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"message": "Internal Server Error",
		"error":   err.Error(),
	}); jsonErr != nil {
		echologrus.Logger.Error(jsonErr)
	}
}
