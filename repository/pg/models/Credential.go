package models

import (
	"time"
)

// Credential model gorm
type Credential struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint `gorm:"column:userId" sql:"index"`
	User      User
	Email     string     `gorm:"column:email" sql:"index"`
	Password  string     `gorm:"column:password" sql:"index"`
	CreatedAt time.Time  `gorm:"column:CreatedAt" sql:"index"`
	UpdatedAt time.Time  `gorm:"column:UpdatedAt" sql:"index"`
	DeletedAt *time.Time `gorm:"column:DeletedAt" sql:"index"` // *time.Time to support nil on gorm model
}

// TableName describe name of the table
func (Credential) TableName() string {
	return "Credentials"
}
