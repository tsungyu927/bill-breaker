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

// @Summary Get book list
// @Description get book list with user_id
// @Tags books
// @Accept json
// @Produce json
// @Param X-User-ID header string true "User ID for authentication"
// @Param book body validators.GetBookListRequest true "Get book list with user_id"
// @Success 200 {object} utils.APIResponse{data=[]models.BookModel} "success"
// @Failure 400 {object} utils.APIResponse "bad request"
// @Failure 404 {object} utils.APIResponse "not found"
// @Failure 500 {object} utils.APIResponse "internal server error"
// @Router /api/v1/book/list [get]
func GetAllBooks(c *gin.Context) {
	userID := c.Request.Header.Get("X-User-ID")

	if err := validators.ValidateGetBookList(validators.GetBookListRequest{ID: userID}); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	var book models.BookModel
	books, err := book.GetBookList(userID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to get book list"))
		return
	}

	response := utils.SuccessResponse(books)
	c.JSON(http.StatusOK, response)
}
