package models

type User struct {
	ID string `json:"id"`
	DeviceId string `json:"device_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	CreatedAt string `json:"created_at"`
	LastModifiedAt string `json:"last_modified_at"`
}

