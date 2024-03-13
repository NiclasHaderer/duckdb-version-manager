package manager

import (
	"duckdb-version-manager/api"
	"duckdb-version-manager/config"
	"duckdb-version-manager/models"
	"duckdb-version-manager/utils"
	"encoding/json"
	"os"
)

var Run VersionManager

func create() (VersionManager, error) {
	// If the config file doesn't exist, create it
	if _, err := os.Stat(config.File); os.IsNotExist(err) {
		err := os.WriteFile(config.File, []byte("{\"localInstallations\": {}}"), 0600)
		if err != nil {
			return nil, err
		}
	}

	// Read the config file
	bytes, err := os.ReadFile(config.File)
	if err != nil {
		return nil, err
	}

	var localConfig models.LocalConfig
	err = json.Unmarshal(bytes, &localConfig)
	if err != nil {
		return nil, err
	}

	return &VersionManagerImpl{
		localConfig: localConfig,
		client:      api.New(),
	}, nil
}

func init() {
	var err error
	Run, err = create()
	if err != nil {
		utils.ExitWithError(err)
	}
}
