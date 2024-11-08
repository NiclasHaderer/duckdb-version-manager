package cmd

import (
	"duckdb-version-manager/cmd/list"
	"duckdb-version-manager/config"
)

func init() {
	rootCmd.Version = config.Version
	rootCmd.AddCommand(defaultCmd)
	rootCmd.AddCommand(updateSelfCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(uninstallSelfCmd)
	rootCmd.AddCommand(uninstallCmd)
	rootCmd.AddCommand(runCmd)

	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(list.LocalCmd)
	listCmd.AddCommand(list.RemoteCmd)
}
