package model

import "gorm.io/gorm"

// 实体

type BannedIpv4 struct {
	gorm.Model
	LeftIp  string `gorm:"type:varchar(32);not null;unique"`
	RightIp string `gorm:"type:varchar(32);not null;unique"`
	User    *User
	UserId  uint `gorm:"not null"`
}

type FailedUser struct {
	gorm.Model
	FailedIp     string `gorm:"type:varchar(32);not null;unique"`
	FailedUser   *User
	FailedUserId uint `gorm:"not null"`
	Trys         uint `gorm:"not null"`
}
