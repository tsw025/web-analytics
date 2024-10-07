package services

import (
	"errors"
	"github.com/hibiken/asynq"
	"github.com/tsw025/web_analytics/internal/echologrus"
	"github.com/tsw025/web_analytics/internal/models"
	"github.com/tsw025/web_analytics/internal/repositories"
	"github.com/tsw025/web_analytics/internal/tasks"
	"gorm.io/gorm"
)

type AnalyseService interface {
	Analyse(url string, user *models.User) (*models.Analytics, error)
}

type analyseService struct {
	websiteRepo  repositories.WebsiteRepository
	analyticRepo repositories.AnalyticsRepository
	asynqClient  *asynq.Client
}

func NewAnalyseService(
	websiteRepo repositories.WebsiteRepository,
	analyticRepo repositories.AnalyticsRepository,
	asynqClient *asynq.Client,
) AnalyseService {
	return &analyseService{
		websiteRepo:  websiteRepo,
		analyticRepo: analyticRepo,
		asynqClient:  asynqClient,
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

	// Create a task to analyze the website
	if analytics.Status == models.StatusPending || analytics.Status == models.StatusFailed {
		if err := s.enqueueAnalyzeTask(url, analytics.ID); err != nil {
			return nil, err
		}
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

func (s *analyseService) enqueueAnalyzeTask(url string, analyticsID uint) error {
	payload := &tasks.AnalyzeWebsitePayload{
		URL:         url,
		AnalyticsID: analyticsID,
	}

	payloadBytes, err := payload.Marshal()
	if err != nil {
		echologrus.Logger.Errorf("Failed to marshal payload: %v", err)
		return err
	}

	task := asynq.NewTask(tasks.TypeAnalyzeWebsite, payloadBytes)

	opts := []asynq.Option{
		asynq.Queue("default"),
		asynq.ProcessIn(0),
	}

	info, err := s.asynqClient.Enqueue(task, opts...)
	if err != nil {
		echologrus.Logger.Errorf("Failed to enqueue task: %v", err)
		return err
	}

	echologrus.Logger.Infof("Task enqueued: %v", info)
	return nil
}
