package validators

type CreateBookRequest struct {
	BookName        string  `json:"book_name" validate:"required,min=1"`
	BookDescription *string `json:"book_description" validate:"omitempty"`
}

type GetBookListRequest struct {
	ID string `uri:"id" validate:"required,uuid4"`
}

type GetBookByIDRequest struct {
	BookID string `uri:"book_id" validate:"required,uuid4"`
}

type JoinBookRequest struct {
	BookID string `json:"book_id" validate:"required,uuid4"`
}

func ValidateCreateBook(data CreateBookRequest) error {
	return validate.Struct(data)
}

func ValidateGetBookList(data GetBookListRequest) error {
	return validate.Struct(data)
}

func ValidateGetBookByID(data GetBookByIDRequest) error {
	return validate.Struct(data)
}

func ValidateJoinBook(data JoinBookRequest) error {
	return validate.Struct(data)
}
