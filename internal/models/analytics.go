package models

import (
	"time"

	"gorm.io/datatypes"
)

type AnalyticsStatus string

const (
	StatusPending    AnalyticsStatus = "pending"
	StatusInProgress AnalyticsStatus = "in_progress"
	StatusCompleted  AnalyticsStatus = "completed"
	StatusFailed     AnalyticsStatus = "failed"
)

type Analytics struct {
	ID        uint            `gorm:"primaryKey"`
	WebsiteID uint            `gorm:"not null;uniqueIndex"`
	Data      datatypes.JSON  `gorm:"type:jsonb"`
	Status    AnalyticsStatus `gorm:"type:analytics_status;not null;default:'pending'"`
	CreatedAt time.Time       `gorm:"type:timestampz"`
	UpdatedAt time.Time       `gorm:"type:timestampz"`

	// Associations
	Website Website `gorm:"foreignKey:WebsiteID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
