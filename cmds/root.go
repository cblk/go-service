package cmds

import (
	"github.com/spf13/cobra"
	"go-service/config"
	"go-service/utils"
	"log"
)

var RootCmd = &cobra.Command{
	Use:     "app",
	Short:   "app server",
	Version: "1.0",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return utils.Try(func() {
			log.Println("Initialize config")

			// Initialization

			// Initialize config
			utils.PanicErr(config.InitConfig(""))

			// Initialize DB
			utils.PanicErr(config.InitDB())
		})
	},
}
