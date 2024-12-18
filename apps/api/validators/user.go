package validators

type CreateUserRequest struct {
	Name  string  `json:"name" validate:"required,min=1"`
	Email *string `json:"email" validate:"omitempty,email"`
}

type GetUserRequest struct {
	ID string `uri:"id" validate:"required,uuid4"`
}

func ValidateCreateUser(data CreateUserRequest) error {
	return validate.Struct(data)
}

func ValidateGetUser(data GetUserRequest) error {
	return validate.Struct(data)
}
