package migrations

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func M20190428(db *gorm.DB) *gormigrate.Gormigrate {
	type AppClient struct {
		ID        uint   `gorm:"primary_key"`
		CreatedAt uint   `json:"created_at" db:"created_at" gorm:"index"` // 每级任务生成时间
		UpdatedAt uint   `json:"updated_at" db:"updated_at" gorm:"index"` // 每级任务完成时间
		AppID     string `json:"app_id" db:"app_id"`                      // 客户端ID
		AppName   string `json:"app_name" db:"app_name"`                  // 客户端名字
		AppAuth   string `json:"app_auth" db:"app_auth"`                  // 客户端授权
	}

	type Task struct {
		ID         uint   `gorm:"primary_key"`
		CreatedAt  uint   `json:"created_at" db:"created_at" gorm:"index"`   // 每级任务生成时间
		FinishedAt uint   `json:"finished_at" db:"finished_at" gorm:"index"` // 每级任务完成时间
		Status     string `json:"status" db:"status"`                        // 任务状态
		ErrType    uint   `json:"err_type" db:"err_type"`                    // 任务错误类型(404，500等)
		Type       uint   `json:"type" db:"type"`                            // 任务类型(article, image)
		AppId      string `json:"app_id" db:"app_id" gorm:"index"`           // APPID
		TaskID     string `json:"task_id" db:"task_id" gorm:"index"`         // 主任务 ID  (uuid.V4().hex())
		Priority   uint8  `json:"priority" db:"priority"`                    // 任务优先度 1-9
		Input      string `json:"input" db:"input"`                          // 任务参数
		RetryNum   string `json:"retry_num" db:"retry_num"`                  // 任务重试次数
		Version    string `json:"version" db:"version"`                      // 任务版本
	}

	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "201904281221_app_client",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&AppClient{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("app_clients").Error
			},
		},
		{
			ID: "201904281221_task",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Task{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("tasks").Error
			},
		},
	})
}
