package router

import (
	"crm/controller"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	user := r.Group("/api/v1/user")
	{
		user.GET("/list", controller.User.List)
		user.POST("/login", controller.User.Login)
	}
	return r
}
