package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Generic wrapper for builds",
	Long: `Detects build system and runs a build. Currently, only supports make`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat("Makefile")
		if err != nil {
			log.Fatal(err)
		}
		out, err := exec.Command("make").Output()
		if err != nil {
			log.Fatalf("%s\n%n\n", out, err)
		}
		log.Printf("%s\n", out)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
