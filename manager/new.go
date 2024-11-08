package manager

import (
	"duckdb-version-manager/api"
	"duckdb-version-manager/config"
	"duckdb-version-manager/manager/migrations"
	"duckdb-version-manager/models"
	"duckdb-version-manager/stacktrace"
	"duckdb-version-manager/utils"
	"encoding/json"
	"os"
)

var Run VersionManager

func init() {
	if tmpErr := config.CreateIfNotExists(); tmpErr != nil {
		utils.ExitWithError(tmpErr)
	}

	if tmpErr := migrations.Run(); tmpErr != nil {
		utils.ExitWithError(tmpErr)
	}

	// Read the config file
	bytes, err := os.ReadFile(config.File)
	if err != nil {
		utils.ExitWithError(stacktrace.Wrap(err))
	}

	var localConfig models.LocalConfig
	err = json.Unmarshal(bytes, &localConfig)
	if err != nil {
		utils.ExitWithError(stacktrace.Wrap(err))
	}

	Run = &versionManagerImpl{
		localConfig: localConfig,
		client:      api.New(),
	}
}
