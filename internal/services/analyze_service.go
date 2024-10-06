package services

import (
	"github.com/tsw025/web_analytics/internal/echologrus"
	"github.com/tsw025/web_analytics/internal/models"
	"github.com/tsw025/web_analytics/internal/repositories"
)

type AnalyseService interface {
	Analyse(url string, user *models.User) (string, error)
}

type analyseService struct {
	websiteRepo repositories.WebsiteRepository
}

func NewAnalyseService(websiteRepo repositories.WebsiteRepository) AnalyseService {
	return &analyseService{
		websiteRepo: websiteRepo,
	}
}

func (s *analyseService) Analyse(url string, user *models.User) (string, error) {
	// Check if the website exists
	_, err := s.websiteRepo.GetByURL(url)
	if err != nil {
		return "", err
	}

	echologrus.Logger.Info(user.Username)

	// Perform the analysis
	return "", nil
}
