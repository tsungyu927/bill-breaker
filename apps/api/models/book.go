package models

import (
	"context"
	"time"

	"github.com/tsungyu927/bill-breaker/api/db"
)

type BookModel struct {
	ID              string    `json:"id"`
	CreatorID       string    `json:"creator_id"`
	BookName        string    `json:"book_name"`
	BookDescription *string   `json:"book_description"`
	CreatedAt       time.Time `json:"create_at"`
	LastModifiedAt  time.Time `json:"last_modified_at"`
}

func (book *BookModel) Create() error {
	query := `INSERT INTO books (creator_id, book_name, book_description) VALUES ($1, $2, $3) RETURNING *`

	return db.GetDB().QueryRow(
		context.Background(),
		query,
		book.CreatorID,
		book.BookName,
		book.BookDescription,
	).Scan(
		&book.ID,
		&book.CreatorID,
		&book.BookName,
		&book.BookDescription,
		&book.CreatedAt,
		&book.LastModifiedAt,
	)
}
