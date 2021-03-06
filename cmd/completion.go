package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const (
	defaultShell = "zsh"
)

var (
	shell string
)

func init() {
	RootCmd.AddCommand(CompletionCmd)

	CompletionCmd.Flags().StringVar(&shell, "shell", defaultShell, "desired shell to generate completions for")
}

// CompletionCmd represents the command for generating completion files for the
// micro cli.
var CompletionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate completion files for the micro cli",
	Run:   completionFunc,
}

func completionFunc(cmd *cobra.Command, args []string) {
	switch strings.ToLower(shell) {
	case "bash":
		RootCmd.GenBashCompletion(os.Stdout)
	case "fish":
		RootCmd.GenFishCompletion(os.Stdout, false)
	case "ps", "powershell", "power_shell":
		RootCmd.GenPowerShellCompletion(os.Stdout)
	case "zsh":
		RootCmd.GenZshCompletion(os.Stdout)
	default:
	}
}
