package service

import (
	"github.com/Chandore998/rental/pkg/users"
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
