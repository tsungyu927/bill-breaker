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

func GetAllBooks(c *gin.Context) {
	data := []gin.H{{"ID": 1, "BookID": "1", "BookName": "Book 1"}, {"ID": 2, "BookId": "2", "BookName": "Book 2"}}

	response := utils.SuccessResponse(data)
	c.JSON(http.StatusOK, response)
}

// @Summary Create new book
// @Description create new book with book info and user_id
// @Tags books
// @Accept json
// @Produce json
// @Param book body validators.CreateBookRequest true "Create new book with info and user_id"
// @Success 200 {object} utils.APIResponse{data=responses.CreateBookResponse} "success"
// @Failure 400 {object} utils.APIResponse "bad request"
// @Failure 404 {object} utils.APIResponse "not found"
// @Failure 500 {object} utils.APIResponse "internal server error"
// @Router /api/v1/book [post]
func CreateNewBook(c *gin.Context) {
	var req validators.CreateBookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid input"))
		return
	}

	if err := validators.ValidateCreateBook(req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	book := models.BookModel{
		CreatorID:       req.CreatorID,
		BookName:        req.BookName,
		BookDescription: req.BookDescription,
	}

	if err := book.Create(); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create user"))
		return
	}

	bookResponse := responses.CreateBookResponse{
		BookModel: book,
	}
	response := utils.SuccessResponse(bookResponse)
	c.JSON(http.StatusOK, response)
}
