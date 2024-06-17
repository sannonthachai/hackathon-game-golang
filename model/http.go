package model

// ServiceCode - service code
const ServiceCode = "71"

// ServiceResponse - response for json service (new standard)
type ServiceResponse struct {
	Status       bool        `json:"status"`
	ServiceCode  string      `json:"service_code"`
	Data         interface{} `json:"data,omitempty"`
	ErrorCode    string      `json:"error_code,omitempty"`
	ErrorTitle   string      `json:"error_title,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
}
