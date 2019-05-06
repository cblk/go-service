package cmds

import (
	"fmt"
	"log"

	"go_service/config"
	"go_service/service"
	"go_service/utils"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "node server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.Try(func() {
			log.Println("start service server")

			conf := config.GetConfig()

			//logy.LoadLogConfig(conf)
			//logy.SetFormat("%L %e %D %T %a %f %S %M")

			app := service.GetHttpApplication()

			utils.PanicErr(app.Run(fmt.Sprintf("%s:%s", conf.GetString("service.host"), conf.GetString("service.port"))))
		})

	},
}
