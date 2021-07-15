package cmds

import (
	"errors"

	"go-service/config"
	"go-service/internal/service/db"
	"go-service/migrate"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
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

	if err := dbi.AutoMigrate(&Migration{}); err != nil {
		panic(err)
	}

	// Register the actual migrations functions below
	migrate.RegisterMigrations()
	migrationInitialized = true
}

func Migrate(dbi *gorm.DB) error {
	if !migrationInitialized {
		return errors.New("migration not initialized")
	}
	return migrate.Migrate(dbi)
}

func Rollback(dbi *gorm.DB) error {
	if !migrationInitialized {
		return errors.New("migration not initialized")
	}
	return migrate.Rollback(dbi)
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
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		logrus.Info("migrate begin")
		// Initialize database
		if err = db.InitDB(config.GetConfig()); err != nil {
			return
		}
		InitMigration(db.GetDB())
		switch _action {
		case "migrate":
			if err = Migrate(db.GetDB()); err != nil {
				logrus.Errorf("Migrate failed, error:%v", err)
				return
			}
			logrus.Info("Migrate succeed!")
			return
		case "rollback":
			if err = Rollback(db.GetDB()); err != nil {
				logrus.Errorf("Rollback failed, error:%v", err)
				return
			}
			logrus.Info("Rollback succeed!")
		default:
			err = errors.New("error action")
			logrus.Errorf("error action:%v", err)
			return
		}
		return
	},
})

func init() {
	RootCmd.AddCommand(MigrateCmd)
}
