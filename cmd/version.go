package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "no version info"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "display version information",
	Long: "display version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("sparky version %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
