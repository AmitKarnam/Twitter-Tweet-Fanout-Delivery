package cmd

import (
	"TweetDelivery/server"

	"github.com/spf13/cobra"
)

var cmdServer = &cobra.Command{
	Use:   "server",
	Short: "Start REST server and initialize required modules.",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return server.RunServer()
	},
}

func init() {
	rootCmd.AddCommand(cmdServer)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmdServerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmdServerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
