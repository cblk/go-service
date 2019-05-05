package migrations

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func M2(db *gorm.DB) *gormigrate.Gormigrate {
	type Task struct {
		ID         uint   `json:"id,omit" gorm:"primary_key"`
		CreatedAt  uint   `json:"created_at,omit" db:"created_at" gorm:"index;not null"`                    // 每级任务生成时间
		FinishedAt uint   `json:"finished_at,omit" db:"finished_at" gorm:"index;not null"`                  // 每级任务完成时间
		Status     string `json:"status,omit" db:"status" gorm:"type:varchar(20);index;not null"`           // 任务状态
		ErrType    uint   `json:"err_type,omit" db:"err_type" gorm:"index;not null"`                        // 任务错误类型(404，500等)
		Type       string `json:"type,omit" db:"type:varchar(20);type;not null"`                            // 任务类型(article, image)
		AppId      string `json:"app_id,omit" db:"app_id" gorm:"type:varchar(100);index;not null"`          // APPID
		TaskID     string `json:"task_id,omit" db:"task_id" gorm:"type:varchar(100);unique_index;not null"` // 主任务 ID  (uuid.V4().hex())
		Priority   uint8  `json:"priority,omit" db:"priority" gorm:"not null"`                              // 任务优先度 1-9
		Input      string `json:"input,omit" db:"input" gorm:"type:text;not null"`                          // 任务参数
		Output     string `json:"output,omit" db:"output" gorm:"type:text;not null"`                        // 任务参数
		RetryNum   int    `json:"retry_num,omit" db:"retry_num" gorm:"not null"`                            // 任务重试次数
		Version    string `json:"version,omit" db:"version" gorm:"type:varchar(20);not null"`               // 任务版本
	}

	type AppClient struct {
		ID        uint   `gorm:"primary_key"`
		CreatedAt uint   `json:"created_at,omit" db:"created_at" gorm:"index"` // 每级任务生成时间
		UpdatedAt uint   `json:"updated_at,omit" db:"updated_at" gorm:"index"` // 每级任务完成时间
		AppID     string `json:"app_id,omit" db:"app_id" gorm:"index"`         // 客户端ID
		AppName   string `json:"app_name,omit" db:"app_name"`                  // 客户端名字
		AppAuth   string `json:"app_auth,omit" db:"app_auth"`                  // 客户端授权
	}

	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "m2_app_client",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&AppClient{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("app_clients").Error
			},
		},
		{
			ID: "m2_task",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Task{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("tasks").Error
			},
		},
	})
}
