package utils

type APIResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message *string     `json:"message"`
}

func SuccessResponse(data interface{}) APIResponse {
	return APIResponse{
		Status:  "success",
		Data:    data,
		Message: nil,
	}
}

func ErrorResponse(message string) APIResponse {
	return APIResponse{
		Status:  "error",
		Data:    nil,
		Message: &message,
	}
}
