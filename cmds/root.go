package cmds

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"go-service/config"
	"go-service/utils"
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
			configPath := "../config"
			err := config.InitConfig(&configPath)

			if err != nil {
				panic(err)
			}

			// Initialize DB
			err = config.InitDB()

			if err != nil {
				panic(err)
			}
		})
	},
}
