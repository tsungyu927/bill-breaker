package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	data := []gin.H{{"ID": 1, "BookID": "1", "BookName": "Book 1"}, {"ID": 2, "BookId": "2", "BookName": "Book 2"}}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": data})
}

func CreateNewBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})	
}
		
