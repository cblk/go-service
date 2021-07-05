package cmds

import (
	"fmt"

	"go-service/api"
	"go-service/config"
	"go-service/internal/service/db"
	"go-service/internal/service/origin"
	"go-service/internal/service/session"
	"go-service/internal/service/storage/oss"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "node server",
	RunE: func(cmd *cobra.Command, args []string) error {
		InitServerFromAppConfig()
		conf := config.GetConfig()
		logrus.Info("start service server")
		app := api.GetHttpApplication(conf)
		address := fmt.Sprintf("%s:%s", conf.Http.Host, conf.Http.Port)
		logrus.Info("server url:" + address)

		return app.Run(address)
	},
}

func InitServerFromAppConfig() {
	appConfig := config.GetConfig()

	// Initialize database
	if err := db.InitDB(appConfig); err != nil {
		panic(err)
	}
	// Initialize session store
	if err := session.InitSessionStore(appConfig); err != nil {
		panic(err)
	}
	// Initialize cors allow origin
	if err := origin.InitOrigin(appConfig); err != nil {
		panic(err)
	}
	// Initialize oss
	if err := oss.InitOss(appConfig); err != nil {
		panic(err)
	}
}

func init() {
	RootCmd.AddCommand(ServerCmd)
}
