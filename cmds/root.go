package cmds

import (
	"github.com/sirupsen/logrus"
	"go_service/config"
	"os"

	"github.com/spf13/cobra"
)

var configPath string
var RootCmd = &cobra.Command{
	Use:     "app",
	Short:   "app server",
	Version: "1.0",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Initialize DB
		return nil
		//return config.InitDB()
	},
}

func Execute() {
	if err := PrepareBaseCmd(RootCmd).Execute(); err != nil {
		panic(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&configPath, "conf", "", "configuration file path")
}

func initConfig() {
	if err := config.InitConfig(configPath); err != nil {
		panic(err)
	} else {
		cfg := config.GetConfig()
		//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
		logrus.SetFormatter(&logrus.TextFormatter{})
		//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
		logrus.SetOutput(os.Stdout)
		//设置最低loglevel
		level, _ := logrus.ParseLevel(cfg.GetString("log.level"))
		logrus.SetLevel(level)
	}
}
