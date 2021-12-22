package config

import (
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	appConfig *AppConfig
)

// InitConfig Init is an exported method that takes the config from the config file
// and unmarshal it into AppConfig struct
func InitConfig(configPath string) error {
	viper.SetConfigType("yml")
	viper.SetConfigName("config")

	if configPath != "" {
		viper.AddConfigPath(configPath)
	} else {
		viper.AddConfigPath("/app/config")
		viper.AddConfigPath("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		logrus.Error("Read config file failed:" + err.Error())
		return err
	}

	appConfig = &AppConfig{}

	if err := viper.Unmarshal(appConfig); err != nil {
		logrus.Error("Parse config file failed:" + err.Error())
		return err
	}

	return nil
}

func GetConfig() *AppConfig {
	return appConfig
}

func WriteConfig() (err error) {
	b, err := json.Marshal(GetConfig())
	if err != nil {
		return
	}
	var m map[string]interface{}
	if err = json.Unmarshal(b, &m); err != nil {
		return
	}
	if err = viper.MergeConfigMap(m); err != nil {
		return
	}
	err = viper.WriteConfig()
	return
}
