package main

import (
	"go-service/cmds"
	"go-service/utils"
)

func main() {
	// Register commands

	rootCmd := cmds.RootCmd
	rootCmd.AddCommand(
		cmds.ServerCmd,
		cmds.MigrateCmd,
		cmds.TestCmd,
	)

	utils.P(utils.Try(func() {
		utils.PanicErr(cmds.PrepareBaseCmd(rootCmd).Execute())
	}))
}
