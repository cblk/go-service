package config

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var config *viper.Viper
var db *gorm.DB

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func InitConfig(configPath *string) error {

	var err error
	v := viper.New()
	v.SetConfigType("yml")
	v.SetConfigName("config")

	if configPath != nil {
		v.AddConfigPath(*configPath)
	} else {
		v.AddConfigPath("config/")
	}

	err = v.ReadInConfig()
	if err != nil {
		return errors.New("error on parsing configuration file")
	}
	config = v

	return nil
}

func GetConfig() *viper.Viper {
	return config
}

func InitDB() error {

	var err error

	config := GetConfig()

	db, err = gorm.Open(config.GetString("db.driver"), config.GetString("db.connection"))

	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
