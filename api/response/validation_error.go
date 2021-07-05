package response

type ValidationErrorResponse struct {
	ErrorResponse
	FieldName string `json:"field_name" description:"The name of the field that fails the validation"`
}

func (ver *ValidationErrorResponse) SetFieldName(fieldName string) {
	ver.FieldName = fieldName
}

func (ver *ValidationErrorResponse) GetFieldName() string {
	return ver.FieldName
}

func NewValidationErrorResponse() *ValidationErrorResponse {
	r := &ValidationErrorResponse{}
	r.SetErrorType("validation_error")
	return r
}
