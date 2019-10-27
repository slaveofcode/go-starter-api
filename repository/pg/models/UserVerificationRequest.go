package models

import (
	"time"
)

// UserVerificationRequest model gorm
type UserVerificationRequest struct {
	ID              uint       `gorm:"primary_key"`
	UserID          uint       `gorm:"column:userId" sql:"index"`
	Type            string     `gorm:"column:type" sql:"index"`
	VerificationKey string     `gorm:"column:verificationKey" sql:"index"`
	VerifiedAt      *time.Time `gorm:"column:verifiedAt" sql:"index"`
	CreatedAt       time.Time  `gorm:"column:CreatedAt" sql:"index"`
	UpdatedAt       time.Time  `gorm:"column:UpdatedAt" sql:"index"`
	DeletedAt       *time.Time `gorm:"column:DeletedAt" sql:"index"` // *time.Time to support nil on gorm model
}

// TableName describe name of the table
func (UserVerificationRequest) TableName() string {
	return "UserVerificationRequests"
}
