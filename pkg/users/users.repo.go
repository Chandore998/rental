package users

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

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

func (user *UsersRepo) IsUsersExist(key string, value string) (bool, error) {
	var userDetail Users

	err := user.DB.Where(key+" = ?", value).First(&userDetail).Error

	fmt.Println(err, "get deeeeeeee")

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		} else {
			return false, nil
		}
	}
	return true, nil
}
