package models

import (
	"encoding/json"
	"strings"
	"time"

	logy "github.com/sirupsen/logrus"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewTask() *Task {
	var _uuid string
	for {
		if _u, err := uuid.NewUUID(); err != nil {
			logy.Errorf("NewTask_NewUUID: %v", err)
			return nil
		} else {
			_uuid = _u.String()
			break
		}
	}

	return &Task{
		CreatedAt: uint(time.Now().Unix()),
		Status:    "pending",
		Type:      "article",
		AppId:     "123456789",
		TaskID:    strings.ReplaceAll(_uuid, "-", ""),
		Priority:  5,
		RetryNum:  5,
		Version:   "v1.0",
	}
}

type Task struct {
	ID         uint   `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt  uint   `json:"created_at,omitempty" db:"created_at" gorm:"index;not null" description:"每级任务生成时间"`
	FinishedAt uint   `json:"finished_at,omitempty" db:"finished_at" gorm:"index;not null" description:"每级任务完成时间"`
	Status     string `json:"status,omitempty" db:"status" gorm:"type:varchar(20);index;not null" enum:"pending,success,error" description:"任务状态"`
	ErrType    uint   `json:"err_type,omitempty" db:"err_type" gorm:"index;not null" enum:"404,500" description:"任务错误类型"`
	Type       string `json:"type,omitempty" db:"type:varchar(20);type;not null" enum:"article,image" description:"任务类型"`
	AppId      string `json:"app_id,omitempty" db:"app_id" gorm:"type:varchar(100);index;not null" description:"APP ID"`
	TaskID     string `json:"task_id,omitempty" db:"task_id" gorm:"type:varchar(100);unique_index;not null" description:"主任务 ID  (uuid.V4().hex())"`
	Priority   uint8  `json:"priority,omitempty" db:"priority" gorm:"not null" description:"任务优先度 1-9"`
	Input      string `json:"input,omitempty" db:"input" gorm:"type:text;not null" description:"任务参数"`
	Output     string `json:"output,omitempty" db:"output" gorm:"type:text;not null" description:"任务参数"`
	RetryNum   int    `json:"retry_num,omitempty" db:"retry_num" gorm:"not null" description:"任务重试次数"`
	Version    string `json:"version,omitempty" db:"version" gorm:"type:varchar(20);not null" description:"任务版本"`
}

func (t *Task) Save(db *gorm.DB) error {
	t.CreatedAt = uint(time.Now().Unix())
	t.Version = "v1.0"
	err := db.Create(t).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *Task) UpdateStatus(db *gorm.DB, taskId, errType, url, status string) error {
	return db.Model(t).Where("task_id = ?", taskId).Updates(
		M{"finished_at": time.Now().Unix(), "output": url, "status": status, "err_type": errType}).Error
}

func (t *Task) Page(db *gorm.DB, taskId, url, status string) error {
	return db.Model(t).Where("task_id = ?", taskId).Updates(
		M{"finished_at": time.Now().Unix(), "output": url, "status": status}).Error
}

func (t *Task) GetTask(db *gorm.DB, taskId string) error {
	return db.Where("task_id = ?", taskId).Find(t).Error
}

func (t *Task) GetTaskStatus(db *gorm.DB, taskId string) error {
	return db.Where("task_id = ?", taskId).Select("status, output").Find(t).Error
}

func (t *Task) Encode() []byte {
	_dt, err := json.Marshal(t)
	if err != nil {
		logy.Errorf("Encode: %v", err)
		return nil
	}
	return _dt
}
