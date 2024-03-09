package utils

import (
	"duckdb-version-manager/config"
	"duckdb-version-manager/models"
	"fmt"
	"os"
)

func GetInstalledVersions() ([]models.Tuple[string, string], error) {
	entries, err := os.ReadDir(config.VersionDir)
	if err != nil {
		return nil, err
	}

	versions := make([]models.Tuple[string, string], 0)
	for _, e := range entries {
		if !e.IsDir() {
			versions = append(versions, models.Tuple[string, string]{First: e.Name(), Second: config.VersionDir + "/" + e.Name()})
		}
	}
	return versions, nil
}

func GetInstalledVersionPath(version string) (*string, error) {
	versions, err := GetInstalledVersions()
	if err != nil {
		return nil, err
	}
	for _, v := range versions {
		if v.First == version {
			return &v.Second, nil
		}
	}
	return nil, fmt.Errorf("version %s not found", version)
}
