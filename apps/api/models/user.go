package models

import (
	"context"
	"time"

	"github.com/tsungyu927/bill-breaker/api/db"
)

type User struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	LastModifiedAt time.Time `json:"last_modified_at"`
}

func (user *User) Create() error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`

	return db.GetDB().QueryRow(context.Background(), query, user.Name, user.Email).Scan(&user.ID)
}

func (user *User) Get(userID string) error {
	query := `SELECT * FROM users WHERE id = $1`

	return db.GetDB().QueryRow(context.Background(), query, userID).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.LastModifiedAt)
}
