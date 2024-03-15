package manager

import (
	"duckdb-version-manager/api"
	"duckdb-version-manager/config"
	"duckdb-version-manager/models"
	"duckdb-version-manager/stacktrace"
	"duckdb-version-manager/utils"
	"encoding/json"
	"os"
)

var Run VersionManager

func create() (VersionManager, stacktrace.Error) {
	// If the config file doesn't exist, create it
	if _, err := os.Stat(config.File); os.IsNotExist(err) {
		err := os.WriteFile(config.File, []byte("{\"localInstallations\": {}}"), 0600)
		if err != nil {
			return nil, stacktrace.Wrap(err)
		}
	}

	// Read the config file
	bytes, err := os.ReadFile(config.File)
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}

	var localConfig models.LocalConfig
	err = json.Unmarshal(bytes, &localConfig)
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}

	return &VersionManagerImpl{
		localConfig: localConfig,
		client:      api.New(),
	}, nil
}

func init() {
	var err stacktrace.Error
	Run, err = create()
	if err != nil {
		utils.ExitWithError(err)
	}
}
