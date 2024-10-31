package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsungyu927/bill-breaker/api/models"
	"github.com/tsungyu927/bill-breaker/api/utils"
	"github.com/tsungyu927/bill-breaker/api/utils/responses"
	"github.com/tsungyu927/bill-breaker/api/validators"
)

// @Summary Create new user
// @Description create new user with new device_id
// @Tags users
// @Accept json
// @Produce json
// @Param user body validators.CreateUserRequest true "Create new user by device_id"
// @Success 200 {object} utils.APIResponse{data=responses.CreateUserResponse} "success"
// @Failure 400 {object} utils.APIResponse "bad request"
// @Failure 404 {object} utils.APIResponse "not found"
// @Failure 500 {object} utils.APIResponse "internal server error"
// @Router /user [post]
func CreateNewUser(c *gin.Context) {
	var req validators.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid input"))	
		return
	}

	if err := validators.ValidateCreateUser(req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	user := models.User{
		DeviceId: req.DeviceId,
		Name: req.Name,
		Email: req.Email,
	}

	if err := user.Create(); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create user"))
		return
	}

	userResponse := responses.CreateUserResponse{
		ID: user.ID,
	}
	response := utils.SuccessResponse(userResponse)
	c.JSON(http.StatusOK, response)
}
