package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsungyu927/bill-breaker/api/utils"
)

func GetAllBooks(c *gin.Context) {
	data := []gin.H{{"ID": 1, "BookID": "1", "BookName": "Book 1"}, {"ID": 2, "BookId": "2", "BookName": "Book 2"}}

	response := utils.SuccessResponse(data)
	c.JSON(http.StatusOK, response)
}

func CreateNewBook(c *gin.Context) {

	response := utils.SuccessResponse("")
	c.JSON(http.StatusOK, response)
}
