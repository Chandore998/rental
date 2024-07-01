package users

import "time"

type Users struct {
	ID                uint    `gorm:"primaryKey" json:"id"`
	FullName          *string ` gorm:"size:200" json:"fullname"`
	Email             string  `gorm:"size:320;unique" json:"email"`
	Password          string  `json:"-"`
	Phone             *string `gorm:"size:15" json:"phone"`
	AboutMe           *string `gorm:"size:500" json:"aboutMe"`
	IsOtpVerified     bool    `gorm:"default:false" json:"isOtpVerified"`
	IsBusinessAccount bool    `gorm:"default:false" json:"isBusinessAccount"`
	Image             *string `json:"image"`
	Zipcode           *string `gorm:"size:6" json:"zipcode"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
