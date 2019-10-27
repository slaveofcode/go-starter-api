package models

import (
	"time"
)

// APIConfig model gorm
type APIConfig struct {
	ID               uint       `gorm:"primary_key"`
	UserID           uint       `gorm:"column:userId" sql:"index"`
	CallLimits       int        `gorm:"column:callLimits" sql:"index"`
	RefreshHourCycle int        `gorm:"column:refreshHourCycle" sql:"index"`
	MaxAPIKeys       int        `gorm:"column:maxApiKeys" sql:"index"`
	CreatedAt        time.Time  `gorm:"column:CreatedAt" sql:"index"`
	UpdatedAt        time.Time  `gorm:"column:UpdatedAt" sql:"index"`
	DeletedAt        *time.Time `gorm:"column:DeletedAt" sql:"index"` // *time.Time to support nil on gorm model
}

// TableName describe name of the table
func (APIConfig) TableName() string {
	return "ApiConfigs"
}
