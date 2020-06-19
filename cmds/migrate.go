package cmds

import (
	"errors"

	logy "github.com/sirupsen/logrus"
	"go_service/config"
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

	// Set default database charset to utf8mb4
	dbi = dbi.Set("gorm:table_options", "CHARSET=utf8mb4")

	// Migration table uses VARCHAR(255) field
	// which cannot be indexed in MySQL before 5.7
	// so we build the migration table using VARCHAR(150)
	// here before the table initialization inside gorm-migrate

	type Migration struct {
		Id string `gorm:"type:varchar(150);primary_key"`
	}

	dbi.AutoMigrate(&Migration{})

	// Register the actual migrations functions below

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
		logy.Info("migrate begin")
		InitMigration(config.GetDB())

		switch _action {
		case "migrate":
			err := Migrate()
			if err != nil {
				logy.Errorf("Migrate failed, error:%v", err)
				return err
			}

			logy.Info("Migrate succeed!")

			return nil
		case "rollback":
			err := Rollback()
			if err != nil {
				logy.Errorf("Rollback failed, error:%v", err)
				return err
			}

			logy.Info("Rollback succeed!")
		default:
			err := errors.New("error action")
			logy.Errorf("error action:%v", err)
			return err
		}

		return nil
	},
})

func init() {
	RootCmd.AddCommand(MigrateCmd)
}
