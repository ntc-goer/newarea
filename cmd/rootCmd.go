package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "",
}

func Execute() error {
	rootCmd.AddCommand(apiCmd())
	rootCmd.AddCommand(dataProcessCmd())
	return rootCmd.Execute()
}
