package request

type FileUploadRequest struct {
	URL        string                 `json:"url"`
	BodyParams map[string]interface{} `json:"body_params"`
}
