package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go-service/utils"
)

var config *viper.Viper
var db *gorm.DB

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func InitConfig(configPath string) error {
	return utils.Try(func() {
		v := viper.New()
		v.SetConfigType("yml")
		v.SetConfigName("config")

		if configPath != "" {
			v.AddConfigPath(configPath)
		} else {
			v.AddConfigPath("config")
		}

		utils.PanicWrap(v.ReadInConfig(), "error on parsing configuration file")

		config = v
	})
}

func GetConfig() *viper.Viper {
	return config
}

func InitDB() error {
	return utils.Try(func() {
		config := GetConfig()

		var err error
		db, err = gorm.Open(config.GetString("db.driver"), config.GetString("db.connection"))
		utils.PanicErr(err)
	})

}

func GetDB() *gorm.DB {
	return db
}
