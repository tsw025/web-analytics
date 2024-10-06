package services

import (
	"errors"
	"github.com/tsw025/web_analytics/internal/handlers"
	"github.com/tsw025/web_analytics/internal/models"
	"github.com/tsw025/web_analytics/internal/repositories"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthService interface {
	Authenticate(username string, password string) (*models.User, error)
	Register(username string, password string) (*models.User, error)
}

type passwordAuthService struct {
	userRepo repositories.UserRepository
}

func NewPasswordAuthService(userRepo repositories.UserRepository) AuthService {
	return &passwordAuthService{
		userRepo: userRepo,
	}
}

func (s *passwordAuthService) Authenticate(username string, password string) (*models.User, error) {
	invalidErrorMessage := "Invalid username or password."
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, handlers.NewDomainError(http.StatusUnauthorized, invalidErrorMessage, errors.New(invalidErrorMessage))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, handlers.NewDomainError(http.StatusUnauthorized, invalidErrorMessage, errors.New(invalidErrorMessage))
	}

	return user, nil
}

func (s *passwordAuthService) Register(username string, password string) (*models.User, error) {
	_, err := s.userRepo.GetByUsername(username)
	if err == nil {
		return nil, handlers.NewDomainError(http.StatusBadRequest, "Username already exists", errors.New("username"))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
