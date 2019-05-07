package main

import (
	"os"

	"go_service/cmds"
)

func main() {
	// Register commands

	rootCmd := cmds.RootCmd
	rootCmd.AddCommand(
		cmds.ServerCmd,
		cmds.MigrateCmd,
	)

	if err := cmds.PrepareBaseCmd(rootCmd).Execute(); err != nil {
		panic(err)
		os.Exit(-1)
	}
}
