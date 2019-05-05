package models

type AppClient struct {
	ID        uint   `gorm:"primary_key"`
	CreatedAt uint   `json:"created_at,omitempty" db:"created_at" gorm:"index"` // 每级任务生成时间
	UpdatedAt uint   `json:"updated_at,omitempty" db:"updated_at" gorm:"index"` // 每级任务完成时间
	AppID     string `json:"app_id,omitempty" db:"app_id" gorm:"index"`                      // 客户端ID
	AppName   string `json:"app_name,omitempty" db:"app_name"`                  // 客户端名字
	AppAuth   string `json:"app_auth,omitempty" db:"app_auth"`                  // 客户端授权
}
