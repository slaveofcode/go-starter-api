package models

import (
	"time"
)

// TeamMemberInvitation model gorm
type TeamMemberInvitation struct {
	ID            uint       `gorm:"primary_key"`
	TeamID        uint       `gorm:"column:teamId" sql:"index"`
	Email         string     `gorm:"column:email" sql:"index"`
	RoleID        uint       `gorm:"column:roleId" sql:"index"`
	InvitationKey string     `gorm:"column:invitationKey" sql:"index"`
	CreatedAt     time.Time  `gorm:"column:CreatedAt" sql:"index"`
	UpdatedAt     time.Time  `gorm:"column:UpdatedAt" sql:"index"`
	DeletedAt     *time.Time `gorm:"column:DeletedAt" sql:"index"` // *time.Time to support nil on gorm model
}

// TableName describe name of the table
func (TeamMemberInvitation) TableName() string {
	return "TeamMemberInvitations"
}
