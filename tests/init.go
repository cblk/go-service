package tests

import (
	"go_service/api"
	"go_service/cmds"
	"go_service/config"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Application *gin.Engine = nil

func init() {
	// Initialize api application to serve API test calls

	testAppConfig := &config.AppConfig{}

	testAppConfig.Environment = config.EnvTest
	testAppConfig.Db.Driver = ""
	testAppConfig.Log.Level = logrus.DebugLevel.String()
	testAppConfig.Http.Host = "127.0.0.1"
	testAppConfig.Http.Port = "8080"

	if err := cmds.InitAllFromAppConfig(testAppConfig); err != nil {
		panic(err)
	}

	Application = api.GetHttpApplication(testAppConfig)
}
