package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tsungyu927/bill-breaker/api/controllers"
)

func UserRegister(router *gin.RouterGroup) {
	router.POST("/user", controllers.CreateNewUser)
	router.GET("/user/:id", controllers.GetUserByUserId)
}
