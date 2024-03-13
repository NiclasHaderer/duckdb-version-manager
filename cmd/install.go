package cmd

import (
	"duckdb-version-manager/manager"
	"duckdb-version-manager/utils"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [version]",
	Short: "Install a specific version of DuckDB",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := manager.Run.InstallVersion(args[0])
		if err != nil {
			utils.ExitWithError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
