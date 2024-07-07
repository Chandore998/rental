package service

import (
	"net/http"

	"github.com/Chandore998/rental/pkg/users"
	"github.com/Chandore998/rental/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UsersService struct {
	Users *users.UsersRepo
}

func NewUsersService(db *gorm.DB) *UsersService {
	return &UsersService{
		Users: users.NewUsersRepo(db),
	}
}

type LoginUser struct {
	Email    string `json:"email" binding:"required";`
	Password string `json:"password" binding:"required"`
}

type SignupUser struct {
	Email           string `json:"email" binding:"required";`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmpassword" binding:"required"`
}

type UpdateUser struct {
	PhoneNumber       *string `json:"phonenumber" `
	FullName          *string `json:"fullname"`
	ZipCode           *string `json: "zipcode"`
	IsBusinessAccount bool    `json:"isbusinessaccount"`
}

func (s *UsersService) login(c *gin.Context) {
	var loginuser LoginUser
	if err := c.ShouldBindJSON(&loginuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDetail, err := s.Users.GetUserByEmail(loginuser.Email)
	if err != nil {
		c.JSON(400, gin.H{"status": http.StatusBadRequest, "error": "Invalid email. or password"})
		return
	}

	isComparePassword := utils.ComparePassword(userDetail.Password, loginuser.Password)
	if !isComparePassword {
		c.JSON(400, gin.H{"status": http.StatusBadRequest, "error": "Password is invaild"})
		return
	}

	c.JSON(200, gin.H{"status": http.StatusOK, "data": userDetail, "error": "null"})
	return

}

func (s *UsersService) signup(c *gin.Context) {
	var signupuser SignupUser

	if err := c.ShouldBindJSON(&signupuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emailExist, err := s.Users.IsUsersExist("email", signupuser.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err})
		return
	}

	if emailExist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Email already exist "})
		return
	}

	if signupuser.Password != signupuser.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is not match confirm password"})
		return
	}

	hashPassword, err := utils.HashPassword(signupuser.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user := users.Users{
		Email:    signupuser.Email,
		Password: signupuser.Password,
	}

	user.Password = hashPassword
	err = db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Users created successfully",
	})

}

func (s *UsersService) updateUser(c *gin.Context) {
	id := c.Param("id")

	var updateuser UpdateUser

	if err := c.ShouldBindJSON(&updateuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := s.Users.IsUsersExist("id", id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err})
		return
	}

	user := users.Users{
		Phone:             updateuser.PhoneNumber,
		IsBusinessAccount: updateuser.IsBusinessAccount,
		Zipcode:           updateuser.ZipCode,
		FullName:          updateuser.FullName,
	}

	_, err = s.Users.UpdateUser(id, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Users updated successfully",
	})

}

func (s *UsersService) getUser(c *gin.Context) {
	id := c.Param("id")

	userDetail, err := s.Users.GetUserInfo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{"status": http.StatusOK, "data": userDetail, "error": "null"})
}
