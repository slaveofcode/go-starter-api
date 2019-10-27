package models

import (
	"time"
)

// Team model gorm
type Team struct {
	ID        uint       `gorm:"primary_key"`
	Name      string     `gorm:"column:name" sql:"index"`
	CreatedAt time.Time  `gorm:"column:CreatedAt" sql:"index"`
	UpdatedAt time.Time  `gorm:"column:UpdatedAt" sql:"index"`
	DeletedAt *time.Time `gorm:"column:DeletedAt" sql:"index"` // *time.Time to support nil on gorm model
}

// TableName describe name of the table
func (Team) TableName() string {
	return "Teams"
}
