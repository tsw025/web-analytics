package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/tsw025/web_analytics/internal/config"
	"time"

	log "github.com/sirupsen/logrus"
)

type AuthTokenService interface {
	GenerateToken(user_id uint) (string, error)
	ValidateToken(token string) bool
}

type authTokenService struct {
	jwtSecret string
}

func NewAuthTokenService(cfg *config.Config) AuthTokenService {
	return &authTokenService{
		jwtSecret: cfg.JWTSecret,
	}
}

func (t *authTokenService) GenerateToken(user_id uint) (string, error) {
	log.Debugf("Generating Tokens")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user_id,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	secretKey := []byte(t.jwtSecret)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", errors.New("Unable to generate tokens, please contact administrator")
	}

	return tokenString, err
}

func (t *authTokenService) ValidateToken(tokenString string) bool {
	log.Debugf("Validating Token")
	secretKey := []byte(t.jwtSecret)
	token, err := jwt.Parse(tokenString, func(tokenString *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		log.Debugf("JWT token parsing error %v", err)
		return false
	}

	if !token.Valid {
		log.Debugf("Invalid JWT token")
		return false
	}
	return true
}
