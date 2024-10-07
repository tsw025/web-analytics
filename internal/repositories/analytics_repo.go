package repositories

import (
	"github.com/tsw025/web_analytics/internal/models"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AnalyticsRepository interface {
	// Add methods specific to the AnalyticsRepository
	GetByWebsiteID(websiteID uint) (*models.Analytics, error)
	UpdateStatus(id uint, status models.AnalyticsStatus) error

	// Embed the BaseRepository
	GetByID(id uint) (*models.Analytics, error)
	Create(user *models.Analytics) error
	Update(user *models.Analytics) error
	Delete(user *models.Analytics) error
	UpdateDataAndStatus(id uint, json datatypes.JSON, completed models.AnalyticsStatus) error
}

type analyticsRepository struct {
	*BaseRepository[models.Analytics]
}

// NewAnalyticsRepository creates a new AnalyticsRepository with the given gorm.DB
func NewAnalyticsRepository(db *gorm.DB) AnalyticsRepository {
	return &analyticsRepository{NewBaseRepository[models.Analytics](db)}
}

// GetByWebsiteID returns analytics by website ID
func (r *analyticsRepository) GetByWebsiteID(websiteID uint) (*models.Analytics, error) {
	var result models.Analytics
	err := r.db.Where("website_id = ?", websiteID).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, err
}

// UpdateStatus updates the status of the analytics
func (r *analyticsRepository) UpdateStatus(id uint, status models.AnalyticsStatus) error {
	analytics, err := r.GetByID(id)
	if err != nil {
		return err
	}

	analytics.Status = status
	return r.Update(analytics)
}

// UpdateDataAndStatus updates the data and status of the analytics
func (r *analyticsRepository) UpdateDataAndStatus(id uint, json datatypes.JSON, completed models.AnalyticsStatus) error {
	analytics, err := r.GetByID(id)
	if err != nil {
		return err
	}

	analytics.Data = json
	analytics.Status = completed
	return r.Update(analytics)
}
