package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harikrishnakreji/GO-GORM-REST/controller"
)

func UserRouter(router *gin.Engine) {
	router.GET("/", controller.UserController)
}
