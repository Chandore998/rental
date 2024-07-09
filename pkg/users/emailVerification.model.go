package users

import "time"

type EmailVerification struct {
	ID                     uint       `gorm:"primaryKey" json:"id"`
	Email                  string     `gorm:"size:320;unique" json:"email"`
	Otp                    string     `gorm:"size:6" json:"otp"`
	IsVerified             *bool      `gorm:"default:false" json:"isVerified"`
	IsOtpVerified          bool       `gorm:"default:false" json:"isOtpVerified"`
	OtpVerificationAttempt int        `gorm:"default:0" json:"otpVerificationAttempt"`
	ExpireDate             *time.Time `gorm:"default:null" json:"expireDate"`
	CreatedAt              time.Time
	UpdatedAt              time.Time
}
