package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RunCmd() *cobra.Command {
	return &cobra.Command{
		Use: "run",

		Run: func(command *cobra.Command, args []string) {
			handleRunCommand(command, args)
		},
	}
}

func handleRunCommand(command *cobra.Command, args []string) {
	err := OpenTerminal()
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	runCmd := RunCmd()
	rootCmd.AddCommand(runCmd)
}
