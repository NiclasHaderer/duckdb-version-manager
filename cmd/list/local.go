package list

import (
	"duckdb-version-manager/client"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var LocalCmd = &cobra.Command{
	Use:   "local",
	Short: "List locally installed DuckDB versions",
	Run: func(cmd *cobra.Command, args []string) {
		releases, err := client.New().ListAllReleases()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Available releases:")
		for _, release := range releases {
			fmt.Printf("  %s\n", release.Version)
		}
	},
}
