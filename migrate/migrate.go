package migrate

import (
	"github.com/cblk/go-service/migrate/migrations"

	"gorm.io/gorm"

	"github.com/go-gormigrate/gormigrate/v2"
)

var migrationList []*gormigrate.Migration

func Migrate(dbi *gorm.DB) error {
	m := gormigrate.New(dbi, gormigrate.DefaultOptions, migrationList)
	return m.Migrate()
}

func Rollback(dbi *gorm.DB) error {
	m := gormigrate.New(dbi, gormigrate.DefaultOptions, migrationList)
	return m.RollbackLast()
}

func RegisterMigrations() {
	migrationList = append(migrationList, migrations.M1)
}
