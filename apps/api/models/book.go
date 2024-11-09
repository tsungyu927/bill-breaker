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

type BookMemberModel struct {
	BookID   string    `json:"book_id"`
	UserID   string    `json:"user_id"`
	JoinedAt time.Time `json:"joined_at"`
}

func (book *BookModel) Create() error {
	query := `INSERT INTO books (creator_id, book_name, book_description) VALUES ($1, $2, $3) RETURNING *`

	err := db.GetDB().QueryRow(
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
	if err != nil {
		return err
	}

	// Add the creator as a member of the book (to book_members table)
	err = addBookMember(book.ID, book.CreatorID)
	if err != nil {
		return err
	}

	return nil
}

func addBookMember(bookID, userID string) error {
	query := `INSERT INTO book_members (book_id, user_id) VALUES ($1, $2)`

	_, err := db.GetDB().Exec(context.Background(), query, bookID, userID)
	return err
}
