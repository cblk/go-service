package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-service/config"
	"go-service/http"
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

			app := http.GetHttpApplication()

			utils.PanicErr(app.Run(fmt.Sprintf("%s:%s", conf.GetString("http.host"), conf.GetString("http.port"))))
		})

	},
}
