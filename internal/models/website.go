package models

import "time"

type Website struct {
	ID        uint      `gorm:"primaryKey"`
	URL       string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"type:timestampz"`
	UpdatedAt time.Time `gorm:"type:timestampz"`

	// Associations
	Users     []User     `gorm:"many2many:user_websites;"`
	Analytics *Analytics `gorm:"foreignKey:WebsiteID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
