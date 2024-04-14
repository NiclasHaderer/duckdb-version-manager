package main

import (
	"duckdb-version-manager/cmd"
	"duckdb-version-manager/manager"
	"duckdb-version-manager/stacktrace"
	"duckdb-version-manager/utils"
)

func main() {
	utils.BeforeErrorExit(func(err stacktrace.Error) {
		manager.Run.ShowUpdateWarning()
	})

	cmd.Execute()
}
