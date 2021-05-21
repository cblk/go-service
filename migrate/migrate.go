package migrate

import (
	logy "github.com/sirupsen/logrus"

	"github.com/go-gormigrate/gormigrate/v2"
)

var migrations []*gormigrate.Gormigrate

func Migrate() error {
	for _, migration := range migrations {
		err := migration.Migrate()
		if err != nil {
			logy.Errorf("Migrate: %v", err)
			return err
		}
	}

	return nil
}

func Rollback() error {
	lastMigration := migrations[len(migrations)-1]
	err := lastMigration.RollbackLast()
	if err != nil {
		logy.Errorf("Migrate: %v", err)
		return err
	}

	return nil
}

func RegisterMigration(migration *gormigrate.Gormigrate) {
	migrations = append(migrations, migration)
}
