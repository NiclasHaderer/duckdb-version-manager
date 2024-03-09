package cmd

import (
	"duckdb-version-manager/utils"
	"github.com/spf13/cobra"
	"os"
	"syscall"
)

var runCmd = &cobra.Command{
	Use:                   "run [version] [duckdb args]",
	Short:                 "Execute a specific version of DuckDB",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := utils.GetInstalledVersionPathOrInstall(args[0])
		if err != nil {
			utils.ExitWithError(err)
		}
		args[0] = *path
		err = syscall.Exec(args[0], args, os.Environ())
		if err != nil {
			utils.ExitWithError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
