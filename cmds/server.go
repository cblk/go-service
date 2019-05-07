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
		logy.SetFormat("%L %e %D %T %a %S %M")

		logy.Info("start service server", nil)

		app := api.GetHttpApplication()

		address := fmt.Sprintf("%s:%s", conf.GetString("http.host"), conf.GetString("http.port"))
		logy.Info("server url:"+address, nil)
		err := app.Run(address)
		if err != nil {
			return err
		}

		return nil
	},
}
