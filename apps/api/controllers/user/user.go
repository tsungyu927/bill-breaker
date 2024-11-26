package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/tsungyu927/bill-breaker/api/models/user"
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
// @Router /api/v1/user [post]
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
		Name:  req.Name,
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

// @Summary Get user
// @Description get user by id
// @Tags users
// @Accept json
// @Produce json
// @Param X-User-ID header string true "User ID for authentication"
// @Success 200 {object} utils.APIResponse{data=models.User} "success"
// @Failure 400 {object} utils.APIResponse "bad request"
// @Failure 404 {object} utils.APIResponse "not found"
// @Failure 500 {object} utils.APIResponse "internal server error"
// @Router /api/v1/user [get]
func GetUserByUserId(c *gin.Context) {
	userID := c.Request.Header.Get("X-User-ID")

	if err := validators.ValidateGetUser(validators.GetUserRequest{ID: userID}); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	var user models.User
	if err := user.Get(userID); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to get user"))
		return
	}

	response := utils.SuccessResponse(user)
	c.JSON(http.StatusOK, response)
}
