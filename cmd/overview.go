package cmd

import (
	"github.com/spf13/cobra"
)

func OverviewCmd() *cobra.Command {
	return &cobra.Command{
		Use: "overview",

		Run: func(command *cobra.Command, args []string) {
			handleOverviewCommand(command, args)
		},
	}
}

func handleOverviewCommand(command *cobra.Command, args []string) {
	showInstanceNames, _ := command.Flags().GetBool("showInstanceNames")
	PrintOverview(showInstanceNames)
}

func init() {
	overviewCommand := OverviewCmd()

	overviewCommand.Flags().BoolP(
		"showInstanceNames",
		"s",
		false,
		"",
	)

	rootCmd.AddCommand(overviewCommand)
}
