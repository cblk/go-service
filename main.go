package main

import (
	"os"
	"go-service/cmds"
	"go-service/utils"
)

func main() {
	// Register commands

	rootCmd := cmds.RootCmd
	rootCmd.AddCommand(
		cmds.VersionCmd,
		cmds.ServerCmd,
		cmds.MigrateCmd,
		cmds.TestCmd,
	)

	utils.PanicErr(cmds.PrepareBaseCmd(rootCmd, "portal",
		os.ExpandEnv("$PWD/kdata")).Execute())
}
