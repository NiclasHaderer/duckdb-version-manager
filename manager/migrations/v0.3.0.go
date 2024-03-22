package migrations

import (
	"duckdb-version-manager/config"
	"github.com/hashicorp/go-version"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	AddMigration(*version.Must(version.NewVersion("0.3.0")), func(localConfig UntypedConfig) (UntypedConfig, error) {
		// Add the new "version" key to the config
		localConfig["version"] = "0.3.0"

		// Iterate over the installedVersions and move the duckdb binary to the new location
		installedVersions := localConfig["localInstallations"].(UntypedConfig)
		for _, installation := range installedVersions {
			installation := installation.(UntypedConfig)
			// Rename the keys in the installation map (change the first letter to lowercase)
			for key, value := range installation {
				delete(installation, key)
				installation[strings.ToLower(key[:1])+key[1:]] = value
			}

			oldInstallationLocation := installation["location"].(string)

			// Move add "duckdb-" to the location if it's not there
			filename := filepath.Base(oldInstallationLocation)
			if !strings.HasPrefix(filename, "duckdb-") {
				newInstallationLocation := filepath.Join(filepath.Dir(oldInstallationLocation), "duckdb-"+filename)
				_ = os.Rename(oldInstallationLocation, newInstallationLocation)
				installation["location"] = newInstallationLocation
			}

		}

		// If the defaultVersion key is present and not empty
		if defaultVersion, ok := localConfig["defaultVersion"]; ok && defaultVersion != "" && defaultVersion != nil {
			// Get the default version file and check where the symlink points to
			symlinkLocation, _ := os.Readlink(config.DefaultDuckdbFile)

			filename := filepath.Base(symlinkLocation)
			if !strings.HasPrefix(filename, "duckdb-") {
				newInstallationLocation := filepath.Join(filepath.Dir(symlinkLocation), "duckdb-"+filename)
				// Change the symlink to point to the new location
				_ = os.Remove(config.DefaultDuckdbFile)
				_ = os.Symlink(newInstallationLocation, config.DefaultDuckdbFile)
			}

		}

		return localConfig, nil
	})
}
