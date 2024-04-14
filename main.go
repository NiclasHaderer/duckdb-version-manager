package main

import (
	"duckdb-version-manager/cmd"
	"duckdb-version-manager/manager"
)

func main() {
	cmd.Execute()
	manager.Run.ShowUpdateWarning()
}
