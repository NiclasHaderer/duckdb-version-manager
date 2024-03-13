package list

import (
	"duckdb-version-manager/api"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var RemoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "List remote DuckDB versions",
	Run: func(cmd *cobra.Command, args []string) {
		releases, err := api.New().ListAllReleases()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Available releases:")
		for _, release := range releases {
			fmt.Printf("  %s\n", release.Version)
		}
	},
}
