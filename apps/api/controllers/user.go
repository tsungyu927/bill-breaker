package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsungyu927/bill-breaker/api/models"
	"github.com/tsungyu927/bill-breaker/api/utils"
	"github.com/tsungyu927/bill-breaker/api/validators"
)

func CreateNewUser(c *gin.Context) {
	var req validators.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errMsg := []string{"Invalid input"}
		errRes := utils.ErrorResponse(errMsg)

		log.Println(err)
		
		c.JSON(http.StatusBadRequest, errRes)	
		return
	}

	if err := validators.ValidateCreateUser(req); err != nil {
		errMsg := []string{err.Error()}
		errRes := utils.ErrorResponse(errMsg)

		log.Println(err)

		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	user := models.User{
		DeviceId: req.DeviceId,
		Name: req.Name,
		Email: req.Email,
	}

	if err := user.Create(); err != nil {
		errMsg := []string{"Failed to create user"}
		errRes := utils.ErrorResponse(errMsg)

		log.Println(err)

		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	
	response := utils.SuccessResponse(gin.H{"id": user.ID})
	c.JSON(http.StatusOK, response)
}

