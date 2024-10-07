package repositories

import (
	"errors"
	"github.com/tsw025/web_analytics/internal/models"
	"gorm.io/gorm"
)

type WebsiteRepository interface {
	// Add methods specific to the WebsiteRepository
	GetByURL(url string) (*models.Website, error)

	// Embed the BaseRepository
	GetByID(id uint) (*models.Website, error)
	GetAll() ([]models.Website, error)
	Create(website *models.Website) error
	Update(website *models.Website) error
	Delete(website *models.Website) error
	GetByUserID(id uint) ([]models.Website, error)
	GetByIDPreloadAnalytics(id uint) (*models.Website, error)
}

type websiteRepository struct {
	*BaseRepository[models.Website]
}

// NewWebsiteRepository creates a new WebsiteRepository with the given gorm.DB
func NewWebsiteRepository(db *gorm.DB) WebsiteRepository {
	return &websiteRepository{NewBaseRepository[models.Website](db)}
}

// GetByURL returns a website by URL
func (r *websiteRepository) GetByURL(url string) (*models.Website, error) {
	var result models.Website
	err := r.db.First(&result, "url = ?", url).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Handle "record not found" case
		return nil, nil // or return a specific error or message
	}

	return &result, err
}

// GetByUserID returns all websites by user ID
func (r *websiteRepository) GetByUserID(id uint) ([]models.Website, error) {
	var result []models.Website
	err := r.db.Find(&result, "user_id = ?", id).Error
	return result, err
}

// GetByIDPreloadAnalytics returns a website by ID with preloaded analytics
func (r *websiteRepository) GetByIDPreloadAnalytics(id uint) (*models.Website, error) {
	var result models.Website
	err := r.db.Preload("Analytics").First(&result, id).Error
	return &result, err
}
