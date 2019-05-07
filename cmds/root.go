package cmds

import (
	"go_service/config"
	"go_service/library/logy"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "app",
	Short:   "app server",
	Version: "1.0",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		// Initialization

		// Initialize config
		err := config.InitConfig("")
		if err != nil {
			return err
		}

		logy.LoadLogConfig(config.GetConfig())
		logy.SetFormat("%L %e %D %T %a %M")

		logy.Info("Initialize config", nil)

		// Initialize DB
		err = config.InitDB()
		if err != nil {

			return err
		}

		return nil
	},
}
