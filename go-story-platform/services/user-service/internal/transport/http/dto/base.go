package dto

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` //interface giúp chứa bất kỳ dto nào
}
