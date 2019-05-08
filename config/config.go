package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var config *viper.Viper
var db *gorm.DB

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func InitConfig(configPath string) error {
	v := viper.New()
	v.SetConfigType("yml")
	v.SetConfigName("config")

	if configPath != "" {
		v.AddConfigPath(configPath)
	} else {
		v.AddConfigPath("/app/config")
		v.AddConfigPath("config")
	}

	err := v.ReadInConfig()
	if err != nil {
		log.Println("InitConfig Failed:"+err.Error())
		panic(err.Error())
		return err
	}

	config = v

	return nil
}

func GetConfig() *viper.Viper {
	return config
}

func InitDB() error {
	config := GetConfig()

	var err error
	db, err = gorm.Open(config.GetString("db.driver"), config.GetString("db.connection"))
	if err != nil {
		log.Println("InitDB Failed:"+err.Error())
		panic(err)
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
