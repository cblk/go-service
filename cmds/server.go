package cmds

import (
	"github.com/spf13/cobra"
	"go-service/config"
	"go-service/http/app"
	"log"
)

var ServerCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "node server",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("start http server")

		conf := config.GetConfig()

		app := app.GetHttpApplication()

		return app.Run(conf.GetString("http.host"), ":", conf.GetString("http.port"))
	},
}
