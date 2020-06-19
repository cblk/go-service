package cmds

import (
	"fmt"

	logy "github.com/sirupsen/logrus"
	"go_service/api"
	"go_service/config"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "node server",
	RunE: func(cmd *cobra.Command, args []string) error {
		conf := config.GetConfig()
		logy.Info("start service server")

		app := api.GetHttpApplication(conf)
		address := fmt.Sprintf("%s:%s", conf.Http.Host, conf.Http.Port)

		logy.Info("server url:" + address)
		return app.Run(address)
	},
}

func init() {
	RootCmd.AddCommand(ServerCmd)
}
