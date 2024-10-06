package services

import (
	"errors"
	"github.com/tsw025/web_analytics/internal/echologrus"
	"github.com/tsw025/web_analytics/internal/models"
	"github.com/tsw025/web_analytics/internal/repositories"
	"gorm.io/gorm"
)

type AnalyseService interface {
	Analyse(url string, user *models.User) (*models.Analytics, error)
}

type analyseService struct {
	websiteRepo  repositories.WebsiteRepository
	analyticRepo repositories.AnalyticsRepository
}

func NewAnalyseService(websiteRepo repositories.WebsiteRepository, analyticRepo repositories.AnalyticsRepository) AnalyseService {
	return &analyseService{
		websiteRepo:  websiteRepo,
		analyticRepo: analyticRepo,
	}
}

func (s *analyseService) Analyse(url string, user *models.User) (*models.Analytics, error) {
	website, err := s.getOrCreateWebsite(url, user)
	if err != nil {
		return nil, err
	}

	analytics, err := s.getOrCreateAnalytics(website.ID)
	if err != nil {
		return nil, err
	}

	return analytics, nil
}

func (s *analyseService) getOrCreateWebsite(url string, user *models.User) (*models.Website, error) {
	website, err := s.websiteRepo.GetByURL(url)
	if err != nil {
		return nil, err
	}

	if website == nil {
		website = &models.Website{URL: url}
		if err := s.websiteRepo.Create(website); err != nil {
			return nil, err
		}
		website.Users = append(website.Users, *user)
		if err := s.websiteRepo.Update(website); err != nil {
			return nil, err
		}
	} else {
		echologrus.Logger.Infof("Website found: %s", website.URL)
	}

	return website, nil
}

func (s *analyseService) getOrCreateAnalytics(websiteID uint) (*models.Analytics, error) {
	analytics, err := s.analyticRepo.GetByWebsiteID(websiteID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			analytics = &models.Analytics{
				WebsiteID: websiteID,
				Status:    models.StatusPending,
			}
			if err := s.analyticRepo.Create(analytics); err != nil {
				return nil, err
			}
			echologrus.Logger.Infof("Analytics created for website ID: %d", websiteID)
		} else {
			return nil, err
		}
	}
	echologrus.Logger.Infof("Analytics found for website ID: %d", websiteID)
	return analytics, nil
}
