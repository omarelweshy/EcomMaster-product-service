package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	// userRepository := &repository.UserRepository{DB: db}
	// userService := &service.UserService{Repo: *userRepository}
	// userHandler := &handler.UserHandler{UserService: userService}
	//
	r := gin.Default()

	// r.Use(middleware.Logging())
	// r.Use(middleware.ErrorHandler())

	return r
}
