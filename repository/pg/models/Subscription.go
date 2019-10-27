package models

import (
	"time"
)

// Subscription model gorm
type Subscription struct {
	ID          uint       `gorm:"primary_key"`
	UserID      uint       `gorm:"column:userId" sql:"index"`
	PlanType    string     `gorm:"column:planType" sql:"index"`
	EndOfPlanAt time.Time  `gorm:"column:endOfPlanAt" sql:"index"`
	IsRecurring bool       `gorm:"column:isRecurring" sql:"index"`
	CreatedAt   time.Time  `gorm:"column:CreatedAt" sql:"index"`
	UpdatedAt   time.Time  `gorm:"column:UpdatedAt" sql:"index"`
	DeletedAt   *time.Time `gorm:"column:DeletedAt" sql:"index"` // *time.Time to support nil on gorm model
}

// TableName describe name of the table
func (Subscription) TableName() string {
	return "Subscriptions"
}
