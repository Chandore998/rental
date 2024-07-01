package users

import "gorm.io/gorm"

type UsersRepo struct {
	DB *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{DB: db}
}

func (user *UsersRepo) GetUserByEmail(email string) (Users, error) {
	var userDetail Users
	err := user.DB.Where("email = ?", email).First(&userDetail).Error
	return userDetail, err

}
