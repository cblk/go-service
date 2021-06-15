package cmds

import (
	"fmt"

	"go_service/api"
	"go_service/config"
	"go_service/config/origin"

	logy "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "node server",
	RunE: func(cmd *cobra.Command, args []string) error {
		InitServerFromAppConfig()
		conf := config.GetConfig()
		logy.Info("start service server")
		app := api.GetHttpApplication(conf)
		address := fmt.Sprintf("%s:%s", conf.Http.Host, conf.Http.Port)
		logy.Info("server url:" + address)

		return app.Run(address)
	},
}

func InitServerFromAppConfig() {
	appConfig := config.GetConfig()

	// Initialize cors allow origin
	if err := origin.InitOrigin(appConfig); err != nil {
		panic(err)
	}
}

func init() {
	RootCmd.AddCommand(ServerCmd)
}
