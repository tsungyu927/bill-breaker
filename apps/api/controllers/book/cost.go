package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/tsungyu927/bill-breaker/api/models/book"
	"github.com/tsungyu927/bill-breaker/api/utils"
	"github.com/tsungyu927/bill-breaker/api/utils/responses"
	"github.com/tsungyu927/bill-breaker/api/validators"
)

// @Summary Create new cost
// @Description create new cost with cost information
// @Tags costs
// @Accept json
// @Produce json
// @Param X-User-ID header string true "User ID for authentication"
// @Param cost body validators.CreateCostRequest true "Create new cost with cost information"
// @Success 200 {object} utils.APIResponse{data=responses.CreateCostResponse} "success"
// @Failure 400 {object} utils.APIResponse "bad request"
// @Failure 404 {object} utils.APIResponse "not found"
// @Failure 500 {object} utils.APIResponse "internal server error"
// @Router /api/v1/book/cost [post]
func CreateNewCost(c *gin.Context) {
	creatorID := c.Request.Header.Get("X-User-ID")

	if creatorID == "" {
		log.Println("Missing creator ID in header")
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Missing User ID"))
		return
	}

	var req validators.CreateCostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid input"))
		return
	}

	if err := validators.ValidateCreateCost(req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	cost := models.CostRecordModel{
		BookID:      req.BookID,
		Amount:      req.Amount,
		Description: req.Description,
		CreatorID:   creatorID,
		Currency:    req.Currency,
	}

	err := cost.CreateCost(req.Payers, req.Sharers)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to add cost"))
		return
	}

	response := utils.SuccessResponse(responses.CreateCostResponse{ID: cost.ID})
	c.JSON(http.StatusOK, response)
}

// @Summary Get cost list
// @Description get cost list by book_id
// @Tags costs
// @Accept json
// @Produce json
// @Param X-User-ID header string true "User ID for authentication"
// @Param book_id path string true "Book ID"
// @Success 200 {object} utils.APIResponse{data=[]models.CostRecordModel} "success"
// @Failure 400 {object} utils.APIResponse "bad request"
// @Failure 404 {object} utils.APIResponse "not found"
// @Failure 500 {object} utils.APIResponse "internal server error"
// @Router /api/v1/book/{book_id}/cost/list [get]
func GetCostList(c *gin.Context) {
	userID := c.Request.Header.Get("X-User-ID")
	bookID := c.Param("book_id")

	if userID == "" {
		log.Println("Missing User ID in header")
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Missing User ID"))
		return
	}

	if err := validators.ValidateGetCostList(validators.GetCostListRequest{BookID: bookID}); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	costs, err := models.GetCostListByBookID(bookID, userID)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	response := utils.SuccessResponse(costs)
	c.JSON(http.StatusOK, response)
}
