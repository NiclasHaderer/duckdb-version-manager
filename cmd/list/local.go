package list

import (
	"duckdb-version-manager/models"
	"duckdb-version-manager/utils"
	"fmt"
	"github.com/spf13/cobra"
)

func printVersions(versions []models.Tuple[string, string]) {
	fmt.Println("Installed DuckDB versions:")
	for _, t := range versions {
		padding := 8 - len(fmt.Sprintf("%v", t.First))
		if padding < 1 {
			padding = 1
		}

		fmt.Printf("  %-*v %v\n", padding+len(fmt.Sprintf("%v", t.First)), t.First, t.Second)
	}
}

var LocalCmd = &cobra.Command{
	Use:   "local",
	Short: "List locally installed DuckDB versions",
	Run: func(cmd *cobra.Command, args []string) {
		entries, err := utils.GetInstalledVersions()
		if err != nil {
			utils.ExitWithError(err)
		}
		printVersions(entries)
	},
}
