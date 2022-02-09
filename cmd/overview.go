package cmd

import (
	"github.com/spf13/cobra"
)

func OverviewCmd() *cobra.Command {
	return &cobra.Command{
		Use: "overview",

		Run: func(command *cobra.Command, args []string) {
			handleOverviewCommand(args)
		},
	}
}

func handleOverviewCommand(args []string) {
	printOverview()
}

func init() {
	overviewCommand := OverviewCmd()
	rootCmd.AddCommand(overviewCommand)
}
