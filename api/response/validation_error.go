package response

import "errors"

type ValidationErrorResponse struct {
	ErrorResponse

	Data struct {
		FieldName string `json:"field_name" description:"The name of the field that fails the validation"`
		Message   string `json:"message" description:"The error type of the validation that fails"`
	} `json:"data" description:"The validation error detail"`
}

func (ver *ValidationErrorResponse) SetFieldName(fieldName string) {
	ver.Data.FieldName = fieldName
}

func (ver *ValidationErrorResponse) SetFieldMessage(message string) {
	ver.Data.Message = message
}

func (ver *ValidationErrorResponse) GetFieldName() string {
	return ver.Data.FieldName
}

func (ver *ValidationErrorResponse) GetFieldMessage() string {
	return ver.Data.Message
}

func newValidationErrorResponse() *ValidationErrorResponse {
	r := &ValidationErrorResponse{}
	r.SetErrorType("validation_error")
	return r
}

// NewValidationErrorResponseWithMessage 此方法用于返回客户端参数错误，http状态码为400
// 请不要用此方法返回服务端500错误
// 需要返回500错误时，请直接返回error类型
func NewValidationErrorResponseWithMessage(fieldName, message string) *ValidationErrorResponse {
	r := newValidationErrorResponse()
	r.Message = message
	r.SetFieldName(fieldName)
	r.SetFieldMessage(message)
	return r
}

func NewAccessDeniedResponse() error {
	return errors.New("access denied")
}
