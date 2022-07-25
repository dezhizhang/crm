package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	BaseController
}

var User *UserController

func (this *UserController) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
		"data": nil,
	})
}

func (this *UserController) Login(c *gin.Context) {

}

func init() {
	User = new(UserController)
}
