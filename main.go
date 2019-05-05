package main

import (
	"go-service/cmds"
	"go-service/utils"
	"os"
)

func main() {
	// Register commands

	rootCmd := cmds.RootCmd
	rootCmd.AddCommand(
		cmds.ServerCmd,
		cmds.MigrateCmd,
		cmds.TestCmd,
	)

	if err := cmds.PrepareBaseCmd(rootCmd).Execute(); err != nil {
		utils.P(err)
		os.Exit(-1)
	}
}
