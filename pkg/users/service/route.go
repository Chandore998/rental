package service

import (
	"github.com/Chandore998/rental/pkg/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func Migration(db *gorm.DB) {
	db.AutoMigrate(&users.Users{})
}

func Init(database *gorm.DB, r *gin.Engine) {
	db = database
	Migration(db)

	as := NewUsersService(db)

	userGroup := r.Group("/user")
	{
		userGroup.POST("/signUp", as.signup)
		userGroup.POST("/login", as.login)
		userGroup.PUT("/:id", as.updateUser)
		userGroup.GET("/:id", as.getUser)
	}

}
