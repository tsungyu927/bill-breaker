package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsungyu927/bill-breaker/api/utils"
)

func CreateNewUser(c *gin.Context) {
	response := utils.SuccessResponse("")
	c.JSON(http.StatusOK, response)
}
