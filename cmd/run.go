package cmd

import (
	"duckdb-version-manager/manager"
	"duckdb-version-manager/utils"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:                   "run [version] [duckdb args]",
	Short:                 "Execute a specific version of DuckDB",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	ValidArgsFunction:     manager.Run.LocalVersionList,
	Run: func(cmd *cobra.Command, args []string) {
		err := manager.Run.Run(args[0], args[1:])
		if err != nil {
			utils.ExitWithError(err)
		}
		manager.Run.ShowUpdateWarning()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
