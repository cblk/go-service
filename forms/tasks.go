package forms

import "encoding/json"

type Task struct {
	Type  string `json:"type" binding:"required"`   // 任务类型
	Url   string `json:"url" binding:"required"`    // Url
	AppId string `json:"app_id" binding:"required"` // 客户端ID
}

func (t Task) String() string {
	_dt, _ := json.Marshal(t)
	return string(_dt)
}

type TaskStatus struct {
	ErrType string `json:"err_type" binding:"required"` // 错误类型
	Url     string `json:"url" binding:"required"`      // Url
	Status  string `json:"status" binding:"required"`   // 状态
}

func (t TaskStatus) String() string {
	_dt, _ := json.Marshal(t)
	return string(_dt)
}
