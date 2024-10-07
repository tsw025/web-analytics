package models

type UserWebsite struct {
	UserID    uint `gorm:"primaryKey"`
	WebsiteID uint `gorm:"primaryKey"`
}
