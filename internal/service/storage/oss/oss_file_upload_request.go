package oss

type OssBodyParams struct {
	Key                 string `json:"key"`
	Policy              string `json:"policy"`
	OSSAccessKeyId      string `json:"OSSAccessKeyId"`
	SuccessActionStatus int    `json:"success_action_status"`
	CallBack            string `json:"callback"`
	Signature           string `json:"signature"`
}
