package models

import (
	"context"
	"errors"
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

type CostDetail struct {
	CostRecordModel
	Payers  []CostPayerModel  `json:"payers"`
	Sharers []CostSharerModel `json:"sharers"`
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

func GetCostListByBookID(bookID, userID string) ([]CostRecordModel, error) {
	costs := []CostRecordModel{}

	// Check whether the user is the member of the book
	isMember, err := isBookMember(bookID, userID)

	if isMember == false || err != nil {
		return nil, errors.New("User is not a member of the book")
	}

	// Fetch costs associated with the bookID
	query := `
		SELECT id, book_id, amount, description, creator_id, currency, created_at
		FROM cost_records
		WHERE book_id = $1
		ORDER BY created_at DESC
	`

	rows, err := db.GetDB().Query(context.Background(), query, bookID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cost CostRecordModel

		err := rows.Scan(
			&cost.ID,
			&cost.BookID,
			&cost.Amount,
			&cost.Description,
			&cost.CreatorID,
			&cost.Currency,
			&cost.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		costs = append(costs, cost)
	}

	return costs, nil
}

func GetCostDetail(bookID, costID, userID string) (*CostDetail, error) {
	// Check whether the user is the member of the book
	isMember, err := isBookMember(bookID, userID)

	if isMember == false || err != nil {
		return nil, errors.New("User is not a member of the book")
	}

	query := `
		SELECT
			c.id AS cost_id,
			c.book_id,
			c.amount,
			c.description,
			c.creator_id,
			c.currency,
			c.created_at,
			cp.user_id AS payer_id,
			cp.amount AS payer_amount,
			cs.user_id AS sharer_id,
			cs.share_amount AS sharer_amount
		FROM cost_records c
		LEFT JOIN cost_payers cp ON c.id = cp.cost_id
		LEFT JOIN cost_sharers cs ON c.id = cs.cost_id
		WHERE c.id = $1 AND c.book_id = $2
	`

	rows, err := db.GetDB().Query(context.Background(), query, costID, bookID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse result into GetCostDetailResponse
	cost := &CostDetail{}
	cost.Payers = []CostPayerModel{}
	cost.Sharers = []CostSharerModel{}

	seenPayers := make(map[string]bool)
	seenSharers := make(map[string]bool)

	for rows.Next() {
		var (
			payerUserID, sharerUserID *string
			payerAmount, sharerAmount *float32
		)

		err := rows.Scan(
			&cost.ID,
			&cost.BookID,
			&cost.Amount,
			&cost.Description,
			&cost.CreatorID,
			&cost.Currency,
			&cost.CreatedAt,
			&payerUserID,
			&payerAmount,
			&sharerUserID,
			&sharerAmount,
		)

		if err != nil {
			return nil, err
		}

		if payerUserID != nil && !seenPayers[*payerUserID] {
			// Add unique payer
			cost.Payers = append(cost.Payers, CostPayerModel{
				CostID: cost.ID,
				UserID: *payerUserID,
				Amount: *payerAmount,
			})
			seenPayers[*payerUserID] = true
		}

		if sharerUserID != nil && !seenSharers[*sharerUserID] {
			// Add unique sharer
			cost.Sharers = append(cost.Sharers, CostSharerModel{
				CostID:      cost.ID,
				UserID:      *sharerUserID,
				ShareAmount: *sharerAmount,
			})
			seenSharers[*sharerUserID] = true
		}
	}

	if cost.ID == "" {
		return nil, errors.New("Cost not found")
	}

	return cost, nil
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
