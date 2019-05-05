package migrate

import (
	"go-service/utils"
	"gopkg.in/gormigrate.v1"
)

var migrations []*gormigrate.Gormigrate

func Migrate() error {
	return utils.Try(func() {
		for _, migration := range migrations {
			utils.PanicErr(migration.Migrate())
		}
	})
}

func Rollback() error {
	return utils.Try(func() {
		lastMigration := migrations[len(migrations)-1]
		utils.PanicErr(lastMigration.RollbackLast())
	})
}

func RegisterMigration(migration *gormigrate.Gormigrate) {
	migrations = append(migrations, migration)
}
