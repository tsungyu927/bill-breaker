package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/tsungyu927/bill-breaker/api/controllers/book"
)

func BookRegister(router *gin.RouterGroup) {
	// book
	router.GET("/book/list", controllers.GetAllBooks)
	router.GET("/book/:book_id", controllers.GetBookByID)
	router.POST("/book", controllers.CreateNewBook)
	router.POST("/book/join", controllers.JoinBook)
	router.DELETE("/book/:book_id/leave", controllers.LeaveBook)

	// cost
	router.POST("/book/cost", controllers.CreateNewCost)
}
