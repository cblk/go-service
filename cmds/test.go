package cmds

import (
	"github.com/spf13/cobra"
	"portal/tests"
	"portal/utils"
)

// TestCmd ...
var TestCmd = &cobra.Command{
	Use:     "test",
	Aliases: []string{"t"},
	Short:   "test",
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.Try(func() {
			utils.PanicErr(tests.TestSendTask())
		})
	},
}
