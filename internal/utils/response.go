package utils

// WebResponse defines the structure for standard web responses
type WebResponse struct {
	Code   int32       `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// ErrorResponse defines the structure for error responses
type ErrorResponse struct {
	Code    int32       `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}
