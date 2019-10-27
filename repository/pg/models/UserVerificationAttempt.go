package models

import (
	"time"
)

// UserVerificationAttempt model gorm
type UserVerificationAttempt struct {
	ID                        uint       `gorm:"primary_key"`
	UserID                    uint       `gorm:"column:userId" sql:"index"`
	UserVerificationRequestID uint       `gorm:"column:userVerificationRequestId" sql:"index"`
	UserAgent                 string     `gorm:"column:userAgent" sql:"index"`
	IPAddr                    string     `gorm:"column:ipAddr" sql:"index"`
	CreatedAt                 time.Time  `gorm:"column:CreatedAt" sql:"index"`
	UpdatedAt                 time.Time  `gorm:"column:UpdatedAt" sql:"index"`
	DeletedAt                 *time.Time `gorm:"column:DeletedAt" sql:"index"` // *time.Time to support nil on gorm model
}

// TableName describe name of the table
func (UserVerificationAttempt) TableName() string {
	return "UserVerificationAttempts"
}
