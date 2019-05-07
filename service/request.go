package service

const (
	CONST_Request_Method_GET    = "GET"
	CONST_Request_Method_POST   = "POST"
	CONST_Request_Method_PUT    = "PUT"
	CONST_Request_Method_Delete = "Delete"

	CONST_Content_Type_ApplicationJson       = "application/json"
	CONST_Content_Type_FormData              = "multipart/form-data"
	CONST_Content_Type_ApplicationUrlencoded = "application/x-www-form-urlencoded"

	CONST_Language_none = "none"
	CONST_Language_zh   = "zh"
	CONST_Language_en   = "en"
	CONST_Language_ja   = "ja"
	CONST_Language_ko   = "ko"

	CONST_ResultCode_Success               = 600
	CONST_ResultCode_Unknown_Error         = 699
	CONST_ResultCode_ClientError           = 600
	CONST_ResultCode_InvalidData           = 601
	CONST_ResultCode_ParseJSON_Error       = 602
	CONST_ResultCode_ClientSignature_Error = 603
	CONST_ResultCode_InputParamter_Error   = 604
	CONST_ResultCode_InputParamter_Empty   = 605
	CONST_ResultCode_Server_error          = 606
)
