package cmd

import (
	"duckdb-version-manager/client"
	"github.com/spf13/cobra"
	"log"
)

var installCmd = &cobra.Command{
	Use:   "install [version]",
	Short: "Install a specific version of DuckDB",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]

		possibleVersions, err := client.New().ListAllReleasesDict()
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		if _, ok := possibleVersions[version]; !ok {
			log.Fatalf("Version '%s' not found", version)
		}

	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
