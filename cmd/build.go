package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/op/go-logging"
)

var (
	makeTargets []string
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Generic wrapper for builds",
	Long: `Detects build system and runs a build. Currently, only supports make`,
	Run: func(cmd *cobra.Command, args []string) {
		// Detect Makefile
		_, err := os.Stat("Makefile")
		if err != nil {
			log.Fatal(err)
		}
		for _, target := range makeTargets {
			cmd := exec.Command("make")
			if target != "" {
				cmd = exec.Command("make", target)
			}
			log.Infof("Running make %s...", target)
			cmd.Stdout = NewLogWriter(logging.DEBUG)
			cmd.Stderr = NewLogWriter(logging.ERROR)
			err := cmd.Run()
			if err != nil {
				log.Fatalf("%s\n", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.Flags().StringSliceVarP(&makeTargets, "targets", "t", []string{""}, "Comma-separated ordered list of make targets")
}
