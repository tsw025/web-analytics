package models

import "time"

type Website struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	URL       string    `gorm:"not null" json:"url"`
	CreatedAt time.Time `gorm:"type:timestampz" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestampz" json:"updated_at"`

	// Associations
	Users     []User     `gorm:"many2many:user_websites;" json:"users,omitempty"`
	Analytics *Analytics `gorm:"foreignKey:WebsiteID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"analytics,omitempty"`
}
