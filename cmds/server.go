package cmds

import (
	"github.com/spf13/cobra"
	"go-service/config"
	"go-service/http/app"
	"go-service/utils"
	"log"
)

var ServerCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "node server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.Try(func() {
			log.Println("start http server")

			conf := config.GetConfig()

			app := app.GetHttpApplication()

			return app.Run(config.GetString("http.port"))
		})
	},
}
