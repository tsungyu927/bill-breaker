package utils

type APIResponse struct {
	Status string `json:"status"`
	Data interface{} `json:"data"`
	Messages *[]string `json:"message"`
}

func SuccessResponse(data interface{}) APIResponse {
	return APIResponse{
		Status: "success",
		Data: data,
		Messages: nil,
	}
}

func ErrorResponse(messages []string) APIResponse {
	return APIResponse{
		Status: "error",
		Data: nil,
		Messages: &messages,
	}
}
