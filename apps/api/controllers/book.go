package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	data := []gin.H{{"id": "1", "name": "Book 1"}, {"id": "2", "name": "Book 2"}}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": data})
}

func CreateNewBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})	
}
		
