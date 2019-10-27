package models

import (
	"time"
)

// Invoice model gorm
type Invoice struct {
	ID             uint       `gorm:"primary_key"`
	UserID         uint       `gorm:"column:userId" sql:"index"`
	PlanType       string     `gorm:"column:token" sql:"index"`
	DiscAmount     int        `gorm:"column:currentCalls" sql:"index"`
	DiscPercentage int        `gorm:"column:limits" sql:"index"`
	PaymentType    *time.Time `gorm:"column:lastRefreshedAt" sql:"index"`
	Amount         int        `gorm:"column:refreshHourCycle" sql:"index"`
	Tax            bool       `gorm:"column:isFrozen" sql:"index"`
	TotalAmount    int        `gorm:"column:totalAmount" sql:"index"`
	PaidAt         time.Time  `gorm:"column:paidAt" sql:"index"`
	CreatedAt      time.Time  `gorm:"column:CreatedAt" sql:"index"`
	UpdatedAt      time.Time  `gorm:"column:UpdatedAt" sql:"index"`
	DeletedAt      *time.Time `gorm:"column:DeletedAt" sql:"index"` // *time.Time to support nil on gorm model
}

// TableName describe name of the table
func (Invoice) TableName() string {
	return "Invoices"
}
