package main

import (
	"duckdb-version-manager/cmd"
	"duckdb-version-manager/manager"
)

func main() {
	manager.Run.LocalVersionList(nil, nil, "")
	cmd.Execute()
}
