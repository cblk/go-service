package cmds

import (
	"fmt"

	"go_service/config"
	"go_service/library/logy"
	"go_service/service/api"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "node server",
	RunE: func(cmd *cobra.Command, args []string) error {

		conf := config.GetConfig()

		logy.LoadLogConfig(conf)
		logy.SetFormat("%L %e %D %T %a %M")

		logy.Info("start service server")

		app := api.GetHttpApplication()

		err := app.Run(fmt.Sprintf("%s:%s", conf.GetString("service.host"), conf.GetString("service.port")))
		if err != nil {
			return err
		}

		return nil
	},
}
