package config

import (
	"duckdb-version-manager/models"
	"duckdb-version-manager/stacktrace"
	"encoding/json"
	"os"
)

func saveEmtpyConfig() stacktrace.Error {
	// Just delete the versions directory, and the default duckdb binary file, as there will be no reference to them in the config file
	_ = os.RemoveAll(Dir)
	_ = os.Remove(DefaultDuckdbFile)
	EnsureFoldersExist()

	emptyConfig := models.LocalConfig{
		LocalInstallations: make(models.LocalInstallations),
		DefaultVersion:     nil,
		Version:            Version,
	}
	serialized, _ := json.MarshalIndent(emptyConfig, "", "  ")
	err := os.WriteFile(File, serialized, 0600)
	return stacktrace.Wrap(err)
}

func CreateIfNotExists() stacktrace.Error {
	if _, err := os.Stat(File); os.IsNotExist(err) {
		return saveEmtpyConfig()
	}

	// Check if the config file is empty
	bytes, err := os.ReadFile(File)
	if err != nil {
		return stacktrace.Wrap(err)
	}
	if len(bytes) == 0 {
		return saveEmtpyConfig()
	}

	return nil
}
