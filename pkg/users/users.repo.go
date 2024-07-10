package users

import (
	"errors"

	"gorm.io/gorm"
)

type UsersRepo struct {
	DB *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{DB: db}
}

func (user *UsersRepo) CreateUser(users Users) (*Users, error) {
	err := user.DB.Create(&users).Error

	if err != nil {
		return nil, err
	}

	userDetail, err := user.GetUserByEmail(users.Email)

	if err != nil {
		return nil, err
	}

	return &userDetail, err
}

func (user *UsersRepo) GetUserByEmail(email string) (Users, error) {
	var userDetail Users
	err := user.DB.Where("email = ?", email).First(&userDetail).Error
	return userDetail, err

}

func (user *UsersRepo) IsUsersExist(key string, value string) (bool, error) {
	var userDetail Users

	err := user.DB.Where(key+" = ?", value).First(&userDetail).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		} else {
			return false, nil
		}
	}
	return true, nil
}

func (user *UsersRepo) UpdateUser(id string, userInput Users) (Users, error) {
	err := user.DB.Where("id = ?", id).Updates(userInput).Error
	return userInput, err
}

func (user *UsersRepo) GetUserInfo(id string) (Users, error) {
	var userDetail Users
	err := user.DB.Where("id = ?", id).First(&userDetail).Error
	return userDetail, err
}
