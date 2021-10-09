package cmds

import (
	"os"

	"go-service/config"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var configPath string
var RootCmd = &cobra.Command{
	Use:     "app",
	Short:   "app server",
	Version: "1.0",
}

func Execute() {
	if err := PrepareBaseCmd(RootCmd).Execute(); err != nil {
		panic(err)
	}
}

func init() {
	cobra.OnInitialize(initAllFromConfigFile)
	RootCmd.PersistentFlags().StringVar(&configPath, "conf", "", "configuration file path")
}

func initAllFromConfigFile() {
	// Initialize config from config file
	if err := config.InitConfig(configPath); err != nil {
		panic(err)
	}
	appConfig := config.GetConfig()
	if err := initLogFormat(appConfig); err != nil {
		panic(err)
	}
}

func initLogFormat(appConfig *config.AppConfig) error {
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.TextFormatter{})
	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	logrus.SetOutput(os.Stdout)
	//设置最低loglevel
	level, _ := logrus.ParseLevel(appConfig.Log.Level)
	logrus.SetLevel(level)

	return nil
}
