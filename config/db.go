package config

import (
	"github.com/jinzhu/gorm"
	logy "github.com/sirupsen/logrus"
	"time"
)

var db *gorm.DB

func InitDB(appConfig *AppConfig) (err error) {

	if appConfig.Environment == EnvTest && appConfig.Db.Driver == "" {
		return nil
	}

	db, err = gorm.Open(appConfig.Db.Driver, appConfig.Db.ConnectionString)
	if err != nil {
		logy.Error("InitDB Failed:" + err.Error())
		return err
	}

	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(50)
	db.DB().SetConnMaxLifetime(5 * time.Second)

	return nil
}

func GetDB() *gorm.DB {
	return db
}
