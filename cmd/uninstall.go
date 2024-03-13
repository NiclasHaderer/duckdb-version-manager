package cmd

import (
	"duckdb-version-manager/manager"
	"duckdb-version-manager/utils"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall [version]",
	Args:  cobra.ExactArgs(1),
	Short: "Uninstall a version of DuckDB",
	Run: func(cmd *cobra.Command, args []string) {
		err := manager.Run.UninstallVersion(args[0])
		if err != nil {
			utils.ExitWithError(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
