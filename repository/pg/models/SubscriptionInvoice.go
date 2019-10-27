package models

import (
	"time"
)

// SubscriptionInvoice model gorm
type SubscriptionInvoice struct {
	ID             uint       `gorm:"primary_key"`
	SubscriptionID uint       `gorm:"column:subscriptionId" sql:"index"`
	InvoiceID      uint       `gorm:"column:invoiceId" sql:"index"`
	CreatedAt      time.Time  `gorm:"column:CreatedAt" sql:"index"`
	UpdatedAt      time.Time  `gorm:"column:UpdatedAt" sql:"index"`
	DeletedAt      *time.Time `gorm:"column:DeletedAt" sql:"index"` // *time.Time to support nil on gorm model
}

// TableName describe name of the table
func (SubscriptionInvoice) TableName() string {
	return "SubscriptionInvoices"
}
