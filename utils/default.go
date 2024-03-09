package utils

import (
	"duckdb-version-manager/config"
	"os"
)

func SetDefaultVersion(path string) error {
	// Create a symlink to config.InstallDir with the name "duckdb"
	symlinkPath := config.InstallDir + "/" + config.DuckDBName

	// 1. Check if the symlink already exists
	if _, err := os.Lstat(symlinkPath); err == nil {
		// 2. If it does, remove it
		if err := os.Remove(symlinkPath); err != nil {
			return err
		}
	}

	// 3. Create the symlink
	if err := os.Symlink(path, symlinkPath); err != nil {
		return err
	}

	return nil
}
