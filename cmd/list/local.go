package list

import (
	"duckdb-version-manager/manager"
	"duckdb-version-manager/models"
	"fmt"
	"github.com/spf13/cobra"
)

func printVersions(versions []models.LocalInstallationInfo) {
	fmt.Println("Installed DuckDB versions:")
	for _, t := range versions {
		padding := 8 - len(fmt.Sprintf("%v", t.Version))
		if padding < 1 {
			padding = 1
		}

		fmt.Printf("  %-*v %v\n", padding+len(fmt.Sprintf("%v", t.Version)), t.Version, t.Location)
	}
}

var LocalCmd = &cobra.Command{
	Use:   "local",
	Short: "List locally installed DuckDB versions",
	Run: func(cmd *cobra.Command, args []string) {
		entries := manager.Run.ListInstalledVersions()
		printVersions(entries)
		manager.Run.ShowUpdateWarning()
	},
}
