package cmds

import (
	"github.com/spf13/cobra"
	"log"
	"portal/version"
)

// VersionCmd ...
var VersionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Show version info",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("version")
		log.Println("app version", version.Version)
		log.Println("app commit version", version.CommitVersion)
		log.Println("app build version", version.BuildVersion)
	},
}
