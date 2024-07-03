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

func (s *UsersService) login(c *gin.Context) {
	var loginuser LoginUser
	if err := c.ShouldBindJSON(&loginuser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userDetail, err := s.Users.GetUserByEmail(loginuser.Email)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "error": "Invalid email. or password"})
	}

	c.JSON(200, gin.H{"message": "doe", "doe": userDetail})
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
	}

	if emailExist {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Email already exist "})
	}

	if signupuser.Password != signupuser.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is not match confirm password"})
		return
	}

	hashPassword, err := utils.HashPassword(signupuser.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
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
	return

}
