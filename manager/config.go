package manager

import (
	"duckdb-version-manager/config"
	"duckdb-version-manager/models"
	"duckdb-version-manager/stacktrace"
	"encoding/json"
	"os"
)

func saveEmtpyConfig() stacktrace.Error {
	// Just delete the versions directory, and the default duckdb binary file, as there will be no reference to them in the config file
	_ = os.RemoveAll(config.Dir)
	_ = os.Remove(config.DefaultDuckdbFile)
	config.EnsureFoldersExist()

	emptyConfig := models.LocalConfig{
		LocalInstallations: make(models.LocalInstallations),
		DefaultVersion:     nil,
		Version:            config.Version,
	}
	serialized, _ := json.MarshalIndent(emptyConfig, "", "  ")
	err := os.WriteFile(config.File, serialized, 0600)
	return stacktrace.Wrap(err)
}

func createConfigIfNotExists() stacktrace.Error {
	if _, err := os.Stat(config.File); os.IsNotExist(err) {
		return saveEmtpyConfig()
	}

	// Check if the config file is empty
	bytes, err := os.ReadFile(config.File)
	if err != nil {
		return stacktrace.Wrap(err)
	}
	if len(bytes) == 0 {
		return saveEmtpyConfig()
	}

	return nil
}
