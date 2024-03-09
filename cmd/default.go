package cmd

import (
	"duckdb-version-manager/utils"
	"github.com/spf13/cobra"
)

// defaultCmd represents the default command
var defaultCmd = &cobra.Command{
	Use:   "default [version]",
	Short: "Set a version of DuckDB as default one to use.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path, err := utils.GetInstalledVersionPathOrInstall(args[0])
		if err != nil {
			utils.ExitWithError(err)
		}

		// Create a symlink to .local/bin/duckdb
		err = utils.SetDefaultVersion(*path)
		if err != nil {
			utils.ExitWithError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(defaultCmd)
}
