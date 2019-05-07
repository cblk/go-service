package service

import (
	"net/http"

	"go_service/library/logy"
	"go_service/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type ResponseBase struct {
	ResultCode int         `json:"result_code"`
	ResultMsg  string      `json:"result_msg,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

type Message struct {
	Zh string `json:"zh"`
	En string `json:"en"`
	Kr string `json:"kr"`
}

type ResponseBaseWithMsg struct {
	ResponseBase
	Message *Message `json:"error_message"`
}

var ResultMap map[int]string
var ResultMapObj map[int]*Message

func init() {
}

func Error(cg *gin.Context, resultcode int) *ResponseBase {
	cg.Status(http.StatusBadRequest)
	resultResponse := &ResponseBase{}

	//
	resultResponse.ResultCode = resultcode
	_, ok := ResultMap[resultcode]
	if !ok {
		resultResponse.ResultCode = CONST_ResultCode_Unknown_Error
		resultResponse.ResultMsg = ResultMap[resultResponse.ResultCode]
	}
	resultResponse.Data = nil

	resp, _ := utils.StructToString(resultResponse)
	logy.Error("Response:%v", resp)

	rand := render.JSON{Data: resultResponse}
	if err := rand.Render(cg.Writer); err != nil {
		panic(err)
	}
	return resultResponse
}

func ErrorMsg(cg *gin.Context, resultCode int) *ResponseBaseWithMsg {
	cg.Status(http.StatusBadRequest)
	resultResponse := &ResponseBaseWithMsg{}

	msg := ResultMapObj[resultCode].En
	//
	resultResponse.ResultCode = resultCode
	resultResponse.ResultMsg = msg
	_, ok := ResultMap[resultCode]
	if !ok {
		resultResponse.ResultCode = CONST_ResultCode_Unknown_Error
		resultResponse.ResultMsg = msg
	}
	resultResponse.Data = nil
	resultResponse.Message = ResultMapObj[resultCode]

	resp, _ := utils.StructToString(resultResponse)
	logy.Error("Response:%v", resp)

	rand := render.JSON{Data: resultResponse}
	if err := rand.Render(cg.Writer); err != nil {
		panic(err)
	}
	return resultResponse
}

/*
httpStatus : 200 series,
200	OK	请求成功接收并处理，一般响应中都会有 body
201	Created	请求已完成，并导致了一个或者多个资源被创建，最常用在 POST 创建资源的时候
202	Accepted	请求已经接收并开始处理，但是处理还没有完成。一般用在异步处理的情况，响应 body 中应该告诉客户端去哪里查看任务的状态

http.StatusOK                   = 200 // RFC 7231, 6.3.1
http.StatusCreated              = 201 // RFC 7231, 6.3.2
http.StatusAccepted             = 202 // RFC 7231, 6.3.3
*/
func Success(cg *gin.Context, httpStatus int, data interface{}) {
	cg.Status(httpStatus)

	resultResponse := &ResponseBaseWithMsg{}
	resultResponse.ResultCode = CONST_ResultCode_Success
	resultResponse.ResultMsg = "success"
	resultResponse.Data = data

	rand := render.JSON{Data: resultResponse}
	if err := rand.Render(cg.Writer); err != nil {
		panic(err)
	}
}
