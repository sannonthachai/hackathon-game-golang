package util

import "gitlab.com/sannonthachai/find-the-hidden-backend/model"

// CreateSuccessResponse - create success response
func CreateSuccessResponse(data interface{}) model.ServiceResponse {
	return model.ServiceResponse{
		Status:      true,
		ServiceCode: model.ServiceCode,
		Data:        data,
	}
}

// CreateErrorResponse - create fail response
func CreateErrorResponse(data interface{}, code string, title string, msg string) model.ServiceResponse {
	return model.ServiceResponse{
		Status:       false,
		ServiceCode:  model.ServiceCode,
		Data:         data,
		ErrorCode:    code,
		ErrorTitle:   title,
		ErrorMessage: msg,
	}
}
