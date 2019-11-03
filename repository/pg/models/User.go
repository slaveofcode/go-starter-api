package models

import (
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// User model gorm
type User struct {
	ID             uint `gorm:"primary_key"`
	Name           string
	City           string
	Country        string
	AvatarImgURL   string     `gorm:"column:avatarImgURL"`
	LastLoginAt    *time.Time `gorm:"column:lastLoginAt" sql:"index"`
	BlockedAt      *time.Time `gorm:"column:blockedAt" sql:"index"`
	VerifiedAt     *time.Time `gorm:"column:verifiedAt" sql:"index"`
	Timezone       string
	TimezoneOffset string     `gorm:"column:timezoneOffset" sql:"index"`
	CreatedAt      time.Time  `gorm:"column:CreatedAt" sql:"index"`
	UpdatedAt      time.Time  `gorm:"column:UpdatedAt" sql:"index"`
	DeletedAt      *time.Time `gorm:"column:DeletedAt" sql:"index"` // *time.Tim to support nil on gorm model
}

// TableName describe name of the table
func (User) TableName() string {
	return "Users"
}

func (User) convertNonEmptyTime(t *time.Time, loc *time.Location) *time.Time {
	if t != nil {
		theTimeVal := *t
		t := theTimeVal.In(loc)
		return &t
	}
	return t
}

// AfterFind hook values after find operation
func (u User) AfterFind() (err error) {
	tz, err := strconv.ParseInt(u.TimezoneOffset, 10, 64)
	if err != nil {
		log.Error(err.Error())
		log.Errorf("parsing timezone failed!")
		return
	}

	loc := time.FixedZone("", int(tz))

	u.LastLoginAt = u.convertNonEmptyTime(u.LastLoginAt, loc)
	u.BlockedAt = u.convertNonEmptyTime(u.BlockedAt, loc)
	u.VerifiedAt = u.convertNonEmptyTime(u.VerifiedAt, loc)

	u.CreatedAt = u.CreatedAt.In(loc)
	u.UpdatedAt = u.UpdatedAt.In(loc)
	u.DeletedAt = u.convertNonEmptyTime(u.DeletedAt, loc)

	return
}
