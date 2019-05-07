package migrate

import (
	"go_service/library/logy"

	"gopkg.in/gormigrate.v1"
)

var migrations []*gormigrate.Gormigrate

func Migrate() error {
	for _, migration := range migrations {
		err := migration.Migrate()
		if err != nil {
			logy.ErrorW("Migrate", err).Error()
			return err
		}
	}

	return nil
}

func Rollback() error {
	lastMigration := migrations[len(migrations)-1]
	err := lastMigration.RollbackLast()
	if err != nil {
		logy.ErrorW("Migrate", err).Error()
		return err
	}

	return nil
}

func RegisterMigration(migration *gormigrate.Gormigrate) {
	migrations = append(migrations, migration)
}
