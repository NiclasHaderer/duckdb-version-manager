package list

import (
	"duckdb-version-manager/config"
	"duckdb-version-manager/models"
	"duckdb-version-manager/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os"
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
		// List the locally installed DuckDB versions
		entries, err := os.ReadDir(config.VersionDir)
		if err != nil {
			utils.ExitWithError(err)
		}

		versions := make([]models.Tuple[string, string], 0)
		for _, e := range entries {
			if !e.IsDir() {
				versions = append(versions, models.Tuple[string, string]{First: e.Name(), Second: config.VersionDir + "/" + e.Name()})
			}
		}

		printVersions(versions)
	},
}
