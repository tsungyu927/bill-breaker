package validators

import (
	"errors"
	"time"
)

type CostPayerRequest struct {
	UserID string  `json:"user_id" validate:"required,uuid4"`
	Amount float32 `json:"amount" validate:"required,number"`
}

type CostSharerRequest struct {
	UserID      string  `json:"user_id" validate:"required,uuid4"`
	ShareAmount float32 `json:"share_amount" validate:"required,number"`
}

type CreateCostRequest struct {
	BookID      string              `json:"book_id" validate:"required,uuid4"`
	Title       string              `json:"title" validate:"required"`
	Amount      float32             `json:"amount" validate:"required,number"`
	Description *string             `json:"description" validate:"omitempty"`
	Currency    string              `json:"currency" validate:"required"`
	Date        time.Time           `json:"date" validate:"required"`
	Payers      []CostPayerRequest  `json:"payers" validate:"required"`
	Sharers     []CostSharerRequest `json:"sharers" validate:"required"`
}

type GetCostListRequest struct {
	BookID string `uri:"book_id" validate:"required,uuid4"`
}

type GetCostDetailRequest struct {
	BookID string `uri:"book_id" validate:"required,uuid4"`
	CostID string `uri:"cost_id" validate:"required,uuid4"`
}

func ValidateCreateCost(data CreateCostRequest) error {
	return validate.Struct(data)
}

func ValidateCost(request CreateCostRequest) error {
	var totalPayerAmount float32
	for _, payer := range request.Payers {
		totalPayerAmount += payer.Amount
	}

	if totalPayerAmount != request.Amount {
		return errors.New("Total payer amount does not equal the given amount")
	}

	var totalSharerAmount float32
	for _, sharer := range request.Sharers {
		totalSharerAmount += sharer.ShareAmount
	}

	if totalSharerAmount != request.Amount {
		return errors.New("Total sharer amount does not equal the given amount")
	}

	return nil
}

func ValidateGetCostList(data GetCostListRequest) error {
	return validate.Struct(data)
}

func ValidateGetCostDetail(data GetCostDetailRequest) error {
	return validate.Struct(data)
}
