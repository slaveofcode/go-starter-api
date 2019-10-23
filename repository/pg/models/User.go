package models

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// User model gorm
type User struct {
	gorm.Model
	Name           string
	City           string
	Country        string
	AvatarImgURL   string `gorm:"Column:avatarImgUrl"`
	LastLoginAt    time.Time
	BlockedAt      time.Time
	VerifiedAt     time.Time
	Timezone       string
	TimezoneOffset string
}

// TableName describe name of the table
func (User) TableName() string {
	return "Users"
}

// AfterFind hook values after find operation
func (u User) AfterFind() (err error) {
	tz, err := strconv.ParseInt(u.TimezoneOffset, 10, 64)
	if err != nil {
		log.Errorf("parsing timezone failed!")
		return
	}

	loc := time.FixedZone("", int(tz))

	u.LastLoginAt = u.LastLoginAt.In(loc)
	u.BlockedAt = u.BlockedAt.In(loc)
	u.VerifiedAt = u.VerifiedAt.In(loc)

	return
}
