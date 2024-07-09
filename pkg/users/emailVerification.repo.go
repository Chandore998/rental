package users

import (
	"gorm.io/gorm"
)

type EmailVerificationRepo struct {
	DB *gorm.DB
}

func NewEmailVerificationRepo(db *gorm.DB) *EmailVerificationRepo {
	return &EmailVerificationRepo{DB: db}
}

// func (emailVerified *EmailVerificationRepo) GetEmailVerification(email string) {
// 	var emailVerification EmailVerification
// 	err := emailVerified.DB.Where("email = ?", email).First(&emailVerification).Error
// }
