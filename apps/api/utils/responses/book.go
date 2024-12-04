package responses

import "github.com/tsungyu927/bill-breaker/api/models/book"

type CreateBookResponse struct {
	models.BookModel
}

type CombinedBookDetail struct {
	models.BookDetail
	Costs []models.CostRecordModel `json:"costs"`
}
