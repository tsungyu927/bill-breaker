package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tsungyu927/bill-breaker/api/controllers"
)

func BookRegister(router *gin.RouterGroup) {
	router.GET("/books", controllers.GetAllBooks)
	router.POST("/book", controllers.CreateNewBook)
}
