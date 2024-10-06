package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Username     string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	Email        string    `gorm:"size:100"`
	CreatedAt    time.Time `gorm:"type:timestampz"`
	UpdatedAt    time.Time `gorm:"type:timestampz"`

	// Associations
	Websites []*Website `gorm:"many2many:user_websites;"`
}
