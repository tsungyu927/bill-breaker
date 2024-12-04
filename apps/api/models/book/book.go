package models

import (
	"context"
	"encoding/json"
	"errors"
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

type BookMember struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
}

type BookDetail struct {
	BookModel
	Members []BookMember `json:"members"`
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

func (book *BookModel) GetBookList(userID string) ([]BookModel, error) {
	query := `
		SELECT
			b.id,
			b.creator_id,
			b.book_name,
			b.book_description,
			b.created_at,
			b.last_modified_at
		FROM books b
		JOIN book_members bm ON b.id = bm.book_id
		WHERE bm.user_id = $1
	`

	rows, err := db.GetDB().Query(context.Background(), query, userID)
	if err != nil {
		return nil, errors.New("Failed to query books by User ID " + err.Error())
	}
	defer rows.Close()

	var books []BookModel
	for rows.Next() {
		var b BookModel
		err := rows.Scan(
			&b.ID,
			&b.CreatorID,
			&b.BookName,
			&b.BookDescription,
			&b.CreatedAt,
			&b.LastModifiedAt,
		)

		if err != nil {
			return nil, errors.New("Failed to scan book: " + err.Error())
		}

		books = append(books, b)
	}

	return books, nil
}

func (book *BookModel) GetBookByID(userID, bookID string) (*BookDetail, error) {
	query := `
		SELECT
			b.id,
			b.creator_id,
			b.book_name,
			b.book_description,
			b.created_at,
			b.last_modified_at,
			COALESCE(json_agg(json_build_object(
				'user_id', bm.user_id,
				'user_name', u.name
			)) FILTER (WHERE bm.user_id IS NOT NULL), '[]') AS members
		FROM books b
		INNER JOIN book_members bm ON b.id = bm.book_id
		INNER JOIN users u ON bm.user_id = u.id
		WHERE b.id = $1
		AND EXISTS( 
			SELECT 1 FROM book_members
			WHERE book_id = b.id AND user_id = $2
		)		
		GROUP BY b.id
	`

	row := db.GetDB().QueryRow(context.Background(), query, bookID, userID)
	var result BookDetail
	var membersJSON []byte

	err := row.Scan(
		&result.ID,
		&result.CreatorID,
		&result.BookName,
		&result.BookDescription,
		&result.CreatedAt,
		&result.LastModifiedAt,
		&membersJSON,
	)
	if err != nil {
		return nil, errors.New("Failed to get book by User ID and Book ID: " + err.Error())
	}

	// Parse JSON into Members slice
	err = json.Unmarshal(membersJSON, &result.Members)
	if err != nil {
		return nil, errors.New("Failed to parse members: " + err.Error())
	}

	return &result, nil
}

func (book *BookModel) JoinBook(userID, bookID string) error {
	err := addBookMember(bookID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (book *BookModel) LeaveBook(userID, bookID string) error {
	query := `
		WITH removed_rows AS (
			DELETE FROM book_members
			WHERE book_id = $1
			AND user_id = $2
			AND user_id != (SELECT creator_id FROM books WHERE id = $1)
			RETURNING *
		)
		SELECT COUNT(*) AS rows_removed FROM removed_rows
	`

	var rowsRemoved int
	err := db.GetDB().QueryRow(context.Background(), query, bookID, userID).Scan(&rowsRemoved)

	if err != nil {
		return err
	}

	if rowsRemoved == 0 {
		// Check if the user was successfully removed
		return errors.New("Failed to leave the book! You might be the creator or not a member")
	}

	return nil
}

func addBookMember(bookID, userID string) error {
	query := `INSERT INTO book_members (book_id, user_id) VALUES ($1, $2)`

	_, err := db.GetDB().Exec(context.Background(), query, bookID, userID)
	return err
}
