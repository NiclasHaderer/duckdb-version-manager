package cmd

import (
	"duckdb-version-manager/stacktrace"
	"duckdb-version-manager/utils"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "duckman",
	Short: "A version manager for DuckDB",
}
var version string

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	if version == "" {
		utils.ExitWithError(stacktrace.New("Version not set using compile flags. Use -ldflags \"-X 'duckdb-version-manager/cmd.version=1.0.0'\" to set the version."))
	}
	rootCmd.Version = version
}
