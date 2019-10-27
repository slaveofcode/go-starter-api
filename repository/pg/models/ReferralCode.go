package models

import (
	"time"
)

// ReferralCode model gorm
type ReferralCode struct {
	ID        uint       `gorm:"primary_key"`
	UserID    uint       `gorm:"column:userId" sql:"index"`
	Code      string     `gorm:"column:code" sql:"index"`
	CreatedAt time.Time  `gorm:"column:CreatedAt" sql:"index"`
	UpdatedAt time.Time  `gorm:"column:UpdatedAt" sql:"index"`
	DeletedAt *time.Time `gorm:"column:DeletedAt" sql:"index"` // *time.Time to support nil on gorm model
}

// TableName describe name of the table
func (ReferralCode) TableName() string {
	return "ReferralCodes"
}
