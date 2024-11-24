package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/tsungyu927/bill-breaker/api/db"
	"github.com/tsungyu927/bill-breaker/api/validators"
)

type CostRecordModel struct {
	ID          string    `json:"id"`
	BookID      string    `json:"book_id"`
	Amount      float32   `json:"amount"`
	Description *string   `json:"description"`
	CreatorID   string    `json:"creator_id"`
	Currency    string    `json:"currency"`
	CreatedAt   time.Time `json:"created_at"`
}

type CostPayerModel struct {
	CostID string  `json:"cost_id"`
	UserID string  `json:"user_id"`
	Amount float32 `json:"amount"`
}

type CostSharerModel struct {
	CostID      string  `json:"cost_id"`
	UserID      string  `json:"user_id"`
	ShareAmount float32 `json:"share_amount"`
}

func (cost *CostRecordModel) CreateCost(payers []validators.CostPayerRequest, sharers []validators.CostSharerRequest) error {
	// Use Transactions(Begin) to ensure that a group of operations either all succeed or all fail.
	tx, err := db.GetDB().Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	query := `
		INSERT INTO cost_records (book_id, amount, description, creator_id, currency)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`

	err = tx.QueryRow(
		context.Background(),
		query,
		cost.BookID,
		cost.Amount,
		cost.Description,
		cost.CreatorID,
		cost.Currency,
	).Scan(
		&cost.ID,
		&cost.CreatedAt,
	)

	if err != nil {
		return err
	}

	// Insert payers
	if err := insertPayers(tx, cost.ID, payers); err != nil {
		return err
	}

	// Insert sharers
	if err := insertSharers(tx, cost.ID, sharers); err != nil {
		return err
	}

	return tx.Commit(context.Background())
}

func insertPayers(tx pgx.Tx, costID string, payers []validators.CostPayerRequest) error {
	payerRows := make([][]interface{}, len(payers))

	for i, payer := range payers {
		payerRows[i] = []interface{}{costID, payer.UserID, payer.Amount}
	}

	_, err := tx.CopyFrom(
		context.Background(),
		pgx.Identifier{"cost_payers"},
		[]string{"cost_id", "user_id", "amount"},
		pgx.CopyFromRows(payerRows),
	)

	return err
}

func insertSharers(tx pgx.Tx, costID string, sharers []validators.CostSharerRequest) error {
	sharerRows := make([][]interface{}, len(sharers))

	for i, sharer := range sharers {
		sharerRows[i] = []interface{}{costID, sharer.UserID, sharer.ShareAmount}
	}

	_, err := tx.CopyFrom(
		context.Background(),
		pgx.Identifier{"cost_sharers"},
		[]string{"cost_id", "user_id", "share_amount"},
		pgx.CopyFromRows(sharerRows),
	)

	return err
}

func isBookMember(bookID, userID string) (bool, error) {
	var isMember bool

	query := `
		SELECT EXISTS (
			SELECT 1 FROM book_members
			WHERE book_id = $1 AND user_id = $2
		)
	`

	err := db.GetDB().QueryRow(context.Background(), query, bookID, userID).Scan(&isMember)

	return isMember, err
}
