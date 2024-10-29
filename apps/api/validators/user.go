package validators

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type CreateUserRequest struct {
	DeviceId string `json:"device_id" validate:"required,uuid4"`
	Name string `json:"name" validate:"required,min=1"`
	Email string `json:"email" validate:"omitempty,email"`
}

type GetUserRequest struct {
	ID string `json:"id" validate:"required,uuid4"`
}

func ValidateCreateUser(data CreateUserRequest) error {
	return validate.Struct(data)
}

func ValidateGetUser(data GetUserRequest) error {
	return validate.Struct(data)
}
