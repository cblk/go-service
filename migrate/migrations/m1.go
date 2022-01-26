package migrations

import (
	"github.com/cblk/go-service/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var M1 *gormigrate.Migration

func init() {
	type User struct {
		models.Base
		Name string `json:"name" gorm:"type:varchar(191);not null;unique;comment:The user name"`
	}

	M1 = &gormigrate.Migration{
		ID: "m1",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&User{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		},
	}
}
