package migrations

import (
	"duckdb-version-manager/config"
	"duckdb-version-manager/stacktrace"
	"duckdb-version-manager/utils"
	"encoding/json"
	"github.com/hashicorp/go-version"
	"os"
)

func Run() stacktrace.Error {
	bytes, err := os.ReadFile(config.File)
	if err != nil {
		utils.ExitWithError(stacktrace.Wrap(err))
	}

	var untypedConfig UntypedConfig
	err = json.Unmarshal(bytes, &untypedConfig)
	if err != nil {
		return stacktrace.NewF("Could not parse the config file. If that error persists, delete the %s and try again. Note, that you might have to re-set the default duckdb version", config.Dir)
	}

	var currentVersion *version.Version
	if versionString, ok := untypedConfig["version"].(string); ok {
		currentVersion, err = version.NewVersion(versionString)
		if err != nil {
			return stacktrace.NewF("Could not parse duckman version: \"%s\"", versionString)
		}
	} else {
		currentVersion = version.Must(version.NewVersion("0.0.0"))
	}

	// If the config file version is higher than the current version, we have a problem
	if version.Must(version.NewVersion(config.Version)).LessThan(currentVersion) {
		return stacktrace.NewF("The config file is from a newer version of duckman. This should not happen! Please update duckman to the latest version using the install script.")
	}

	for _, migration := range migrations {
		if migration.ToVersion.LessThanOrEqual(currentVersion) {
			continue
		}
		untypedConfig, err = migration.Up(untypedConfig)
		if err != nil {
			return stacktrace.NewF("Migration from version %s to version %s failed: %s", currentVersion.String(), migration.ToVersion.String(), err.Error())
		}
		currentVersion = &migration.ToVersion
	}

	// Save the new config
	serialized, err := json.MarshalIndent(untypedConfig, "", "  ")
	if err != nil {
		return stacktrace.Wrap(err)
	}

	err = os.WriteFile(config.File, serialized, 0600)
	if err != nil {
		return stacktrace.Wrap(err)
	}
	return nil
}
