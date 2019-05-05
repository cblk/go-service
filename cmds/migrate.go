package cmds

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"log"
	"portal/internal/config"
	"portal/migrate"
	"portal/migrate/migrations"
	"portal/utils"
)

var migrationInitialized bool

func InitMigration(dbi *gorm.DB) {

	if migrationInitialized {
		return
	}

	//migrate.RegisterMigration(migrations.M20190428(dbi))
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
		log.Println("migrate")
		cfg := config.DefaultConfig()
		InitMigration(cfg.GetDb())

		switch _action {
		case "migrate":
			utils.PanicErr(Migrate())
			log.Println("Migration succeed!")
		case "rollback":
			utils.PanicErr(Rollback())
			log.Println("Rollback succeed!")
		default:
			return errors.New("error action")
		}

		return nil
	},
})
