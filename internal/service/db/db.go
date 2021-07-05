package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	"go_service/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB(appConfig *config.AppConfig) (err error) {

	if appConfig.Environment == config.EnvTest && appConfig.Db.Driver == "" {
		return nil
	}

	sqlDB, err := sql.Open(appConfig.Db.Driver, appConfig.Db.ConnectionString)
	if err != nil {
		logrus.Error("InitDB sql.Open error:" + err.Error())
		return err
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(5 * time.Second)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,
		},
	)

	db, err = gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}),
		&gorm.Config{Logger: newLogger})
	if err != nil {
		logrus.Error("InitDB Failed:" + err.Error())
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
