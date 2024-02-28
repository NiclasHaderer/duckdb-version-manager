package main

import (
	"duckdb-version-manager/cmd"
	"duckdb-version-manager/utils"
)

// TODO add a update-self command to update the duckdb version manager
func main() {
	utils.GetDeviceInfo()
	cmd.Execute()
}
