package routes

import (
	user "github.com/Cingihimut/goRestApiGin.git/src/modules/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ctx *gin.Context
)

func Api(router *gin.Engine, db *gorm.DB) {
	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userController := user.NewUserController(userService, ctx)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/user", userController.Index)
		v1.GET("/user/:id", userController.GetById)
		v1.POST("/user", userController.Create)
		v1.PATCH("/user/:id", userController.Update)
		v1.DELETE("/user", userController.Delete)
	}
}
