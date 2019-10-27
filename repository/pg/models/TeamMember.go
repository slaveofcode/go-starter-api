package models

import (
	"time"
)

// TeamMember model gorm
type TeamMember struct {
	ID        uint       `gorm:"primary_key"`
	TeamID    uint       `gorm:"column:teamId" sql:"index"`
	UserID    uint       `gorm:"column:userId" sql:"index"`
	RoleID    uint       `gorm:"column:roleId" sql:"index"`
	IsFrozen  bool       `gorm:"column:isFrozen" sql:"index"`
	CreatedAt time.Time  `gorm:"column:CreatedAt" sql:"index"`
	UpdatedAt time.Time  `gorm:"column:UpdatedAt" sql:"index"`
	DeletedAt *time.Time `gorm:"column:DeletedAt" sql:"index"` // *time.Time to support nil on gorm model
}

// TableName describe name of the table
func (TeamMember) TableName() string {
	return "TeamMembers"
}
