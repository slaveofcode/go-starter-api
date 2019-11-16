package models

import (
	"time"
)

// ResetCredential model gorm
type ResetCredential struct {
	ID           uint `gorm:"primary_key"`
	CredentialID uint `gorm:"column:credentialId" sql:"index"`
	Credential   Credential
	ResetToken   string     `gorm:"column:resetToken" sql:"index"`
	ValidUntil   time.Time  `gorm:"column:validUntil" sql:"index"`
	ValidatedAt  *time.Time `gorm:"column:validatedAt" sql:"index"`
	IsExpired    bool       `gorm:"column:isExpired" sql:"index"`
	CreatedAt    time.Time  `gorm:"column:CreatedAt" sql:"index"`
	UpdatedAt    time.Time  `gorm:"column:UpdatedAt" sql:"index"`
	DeletedAt    *time.Time `gorm:"column:DeletedAt" sql:"index"` // *time.Time to support nil on gorm model
}

// TableName describe name of the table
func (ResetCredential) TableName() string {
	return "ResetCredentials"
}
