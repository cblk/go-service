package cmds

import (
	"os"

	"github.com/spf13/cobra"
)

func PrepareBaseCmd(cmd *cobra.Command) Executor {
	cobra.OnInitialize()
	return Executor{Command: cmd, Exit: os.Exit}
}

// Executor wraps the cobra Command with a nicer Execute method
type Executor struct {
	*cobra.Command
	Exit func(int) // this is os.Exit by default, override in tests
}

type ExitCoder interface {
	ExitCode() int
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func (e Executor) Execute() error {
	e.SilenceUsage = true
	e.SilenceErrors = true

	return e.Command.Execute()
}
