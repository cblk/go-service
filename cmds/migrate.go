package cmds

import (
	"errors"

	"go_service/config"
	"go_service/library/logy"
	"go_service/migrate"
	"go_service/migrate/migrations"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

var migrationInitialized bool

func InitMigration(dbi *gorm.DB) {

	if migrationInitialized {
		return
	}

	// migrate.RegisterMigration(migrations.M20190428(dbi))
	migrate.RegisterMigration(migrations.M2(dbi))

	migrationInitialized = true
}

func Migrate() error {

	if !migrationInitialized {
		return errors.New("migration not initialized")
	}

	return migrate.Migrate()
}

func Rollback() error {

	if !migrationInitialized {
		return errors.New("migration not initialized")
	}

	return migrate.Rollback()
}

var _action = "migrate"

func initMigrateCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringVar(&_action, "action", _action, "数据库操作方法(migrate,rollback)")
	return cmd
}

var MigrateCmd = initMigrateCmd(&cobra.Command{
	Use:     "migrate",
	Aliases: []string{"m"},
	Short:   "migrate",
	RunE: func(cmd *cobra.Command, args []string) error {

		logy.LoadLogConfig(config.GetConfig())
		logy.SetFormat("%L %e %D %T %a %M")

		logy.Info("migrate begin", nil)

		InitMigration(config.GetDB())

		switch _action {
		case "migrate":
			err := Migrate()
			if err != nil {
				logy.Error("Migrate failed, error:%v", err)
				return err
			}

			logy.Info("Migrate succeed!", nil)

			return nil
		case "rollback":
			err := Rollback()
			if err != nil {
				logy.Error("Rollback failed, error:%v", err)
				return err
			}

			logy.Info("Rollback succeed!", nil)
		default:
			err := errors.New("error action")
			logy.Error("error action", err)

			return err
		}

		return nil
	},
})
