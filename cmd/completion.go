package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var shellType string

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates shell completion scripts",
	Long: `Supports zsh or bash shell types
To load completion run

. <(sprk --shell [zsh|bash] completion)

To configure your shell to load completions for each session add to your rc file

# ~/.bashrc or ~/.profile or ~/.zshrc
. <(sprk --shell [zsh|bash] completion)
`,
	Run: func(cmd *cobra.Command, args []string) {
		if shellType == "bash" {
			rootCmd.GenBashCompletion(os.Stdout)
		} else if shellType == "zsh" {
			rootCmd.GenZshCompletion(os.Stdout)
		} else {
			fmt.Printf("Shell type must be zsh or bash")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
	completionCmd.Flags().StringVar(&shellType, "shell", "", "Specifiy either zsh or bash as the shell type")
	err := completionCmd.MarkFlagRequired("shell")
	if err != nil{
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}
