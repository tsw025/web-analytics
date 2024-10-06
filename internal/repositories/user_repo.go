package repositories

import (
	"github.com/tsw025/web_analytics/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	// Add methods specific to the UserRepository
	GetByUsername(username string) (*models.User, error)
	AddWebsite(userID uint, websiteID uint) error
	RemoveWebsite(userID uint, websiteID uint) error
	GetWebsites(userID uint) ([]models.Website, error)

	// Embed the BaseRepository
	GetByID(id uint) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(user *models.User) error
}

type userRepository struct {
	*BaseRepository[models.User]
	websiteRepo WebsiteRepository
}

// NewUserRepository creates a new UserRepository with the given gorm.DB
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{NewBaseRepository[models.User](db), NewWebsiteRepository(db)}
}

// GetByUsername returns a user by username
func (r *userRepository) GetByUsername(username string) (*models.User, error) {
	var result models.User
	err := r.db.Where("username = ?", username).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, err
}

// AddWebsite adds a website to the user
func (r *userRepository) AddWebsite(userID uint, websiteID uint) error {
	user, err := r.GetByID(userID)
	if err != nil {
		return err
	}

	website, err := r.websiteRepo.GetByID(websiteID)
	if err != nil {
		return err
	}

	return r.db.Model(user).Association("Websites").Append(website)
}

// RemoveWebsite removes a website from the user
func (r *userRepository) RemoveWebsite(userID uint, websiteID uint) error {
	user, err := r.GetByID(userID)
	if err != nil {
		return err
	}

	website, err := r.websiteRepo.GetByID(websiteID)
	if err != nil {
		return err
	}

	return r.db.Model(user).Association("Websites").Delete(website)
}

// GetWebsites returns all websites for a user
func (r *userRepository) GetWebsites(userID uint) ([]models.Website, error) {
	user, err := r.GetByID(userID)
	if err != nil {
		return nil, err
	}

	var websites []models.Website
	err = r.db.Model(user).Association("Websites").Find(&websites)
	if err != nil {
		return nil, err
	}

	return websites, nil
}
