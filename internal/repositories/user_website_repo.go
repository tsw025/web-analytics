package repositories

import (
	"github.com/tsw025/web_analytics/internal/models"
	"gorm.io/gorm"
)

type WebsiteUserRepository interface {
	GetByUserID(userID uint) ([]models.UserWebsite, error)
	GetByWebsiteIDAndUserID(websiteID, userID uint) (*models.UserWebsite, error)
	GetWebsitesByUserID(id uint) ([]models.Website, error)

	GetByID(id uint) (*models.UserWebsite, error)
	Create(user *models.UserWebsite) error
	Update(user *models.UserWebsite) error
	Delete(user *models.UserWebsite) error
}

type websiteUserRepository struct {
	*BaseRepository[models.UserWebsite]
}

func NewWebsiteUserRepository(db *gorm.DB) WebsiteUserRepository {
	return &websiteUserRepository{NewBaseRepository[models.UserWebsite](db)}
}

func (r *websiteUserRepository) GetByUserID(userID uint) ([]models.UserWebsite, error) {
	var result []models.UserWebsite
	err := r.db.Find(&result, "user_id = ?", userID).Error
	return result, err
}

func (r *websiteUserRepository) GetByWebsiteIDAndUserID(websiteID, userID uint) (*models.UserWebsite, error) {
	var result models.UserWebsite
	err := r.db.First(&result, "website_id = ? AND user_id = ?", websiteID, userID).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *websiteUserRepository) GetWebsitesByUserID(id uint) ([]models.Website, error) {
	var result []models.Website
	err := r.db.Table("websites").
		Joins("JOIN user_websites ON websites.id = user_websites.website_id").
		Where("user_websites.user_id = ?", id).
		Find(&result).Error
	return result, err
}
