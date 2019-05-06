package main

import (
	"os"

	"go_service/cmds"
	"go_service/utils"
)

func main() {
	// Register commands

	rootCmd := cmds.RootCmd
	rootCmd.AddCommand(
		cmds.ServerCmd,
		cmds.MigrateCmd,
	)

	if err := cmds.PrepareBaseCmd(rootCmd).Execute(); err != nil {
		utils.P(err)
		os.Exit(-1)
	}
}
