package services

import (
	"github.com/tsw025/web_analytics/internal/models"
	"github.com/tsw025/web_analytics/internal/repositories"
)

type WebsiteService interface {
	GetWebsites(user *models.User) ([]models.Website, error)
	GetWebsite(id uint) (*models.Website, error)
	UpdateWebsite(id uint, website *models.Website) (*models.Website, error)
}

type websiteService struct {
	websiteRepo repositories.WebsiteRepository
	userWebRepo repositories.WebsiteUserRepository
}

func NewWebsiteService(websiteRepo repositories.WebsiteRepository, userWebRepo repositories.WebsiteUserRepository) WebsiteService {
	return &websiteService{
		websiteRepo: websiteRepo,
		userWebRepo: userWebRepo,
	}
}

func (s *websiteService) GetWebsites(user *models.User) ([]models.Website, error) {
	websites, err := s.userWebRepo.GetWebsitesByUserID(user.ID)
	if err != nil {
		return nil, err
	}
	return websites, nil
}

func (s *websiteService) GetWebsite(id uint) (*models.Website, error) {
	return s.websiteRepo.GetByIDPreloadAnalytics(id)
}

func (s *websiteService) UpdateWebsite(id uint, website *models.Website) (*models.Website, error) {
	// Check if the website exists
	_, err := s.websiteRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	website.ID = id
	err = s.websiteRepo.Update(website)
	if err != nil {
		return nil, err
	}

	return website, nil
}
