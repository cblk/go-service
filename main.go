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

	utils.PanicErr(cmds.PrepareBaseCmd(rootCmd, "portal",
		os.ExpandEnv("$PWD")).Execute())
}
