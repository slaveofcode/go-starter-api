package models

import (
	"time"
)

// APIKey model gorm
type APIKey struct {
	ID               uint       `gorm:"primary_key"`
	UserID           uint       `gorm:"column:userId" sql:"index"`
	Token            string     `gorm:"column:token" sql:"index"`
	CurrentCalls     int        `gorm:"column:currentCalls" sql:"index"`
	Limits           int        `gorm:"column:limits" sql:"index"`
	LastRefreshedAt  *time.Time `gorm:"column:lastRefreshedAt" sql:"index"`
	RefreshHourCycle int        `gorm:"column:refreshHourCycle" sql:"index"`
	IsFrozen         bool       `gorm:"column:isFrozen" sql:"index"`
	CreatedAt        time.Time  `gorm:"column:CreatedAt" sql:"index"`
	UpdatedAt        time.Time  `gorm:"column:UpdatedAt" sql:"index"`
	DeletedAt        *time.Time `gorm:"column:DeletedAt" sql:"index"` // *time.Time to support nil on gorm model
}

// TableName describe name of the table
func (APIKey) TableName() string {
	return "ApiKeys"
}
