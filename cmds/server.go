package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-service/config"
	"go-service/service"
	"go-service/utils"
	"log"
)

var ServerCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "node server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.Try(func() {
			log.Println("start service server")

			conf := config.GetConfig()

			app := service.GetHttpApplication()

			utils.PanicErr(app.Run(fmt.Sprintf("%s:%s", conf.GetString("service.host"), conf.GetString("service.port"))))
		})

	},
}
