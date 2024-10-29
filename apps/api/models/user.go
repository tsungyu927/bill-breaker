package models

import (
	"context"

	"github.com/tsungyu927/bill-breaker/api/db"
)

type User struct {
	ID string `json:"id"`
	DeviceId string `json:"device_id"`
	Name string `json:"name"`
	Email string `json:"email,omitempty"`
	CreatedAt string `json:"created_at"`
	LastModifiedAt string `json:"last_modified_at"`
}

func (user *User) Create() error {
	query := `INSERT INTO users (device_id, name, email) VALUES ($1, $2, $3) RETURNING id`
	
	return db.GetDB().QueryRow(context.Background(), query, user.DeviceId, user.Name, user.Email).Scan(&user.ID)
}

func (user *User) Get() {
	// TODO: Get specific user from DB
}

