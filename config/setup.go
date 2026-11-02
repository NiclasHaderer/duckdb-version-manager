package config

import (
	"duckdb-version-manager/models"
	"duckdb-version-manager/stacktrace"
	"encoding/json"
	"os"
	"path/filepath"
)

func saveEmtpyConfig() stacktrace.Error {
	// Just delete the versions directory, and the default duckdb binary file, as there will be no reference to them in the config file
	_ = os.RemoveAll(Dir)
	_ = os.Remove(DefaultDuckdbFile)
	EnsureFoldersExist()

	return Save(models.LocalConfig{
		LocalInstallations: make(models.LocalInstallations),
		DefaultVersion:     nil,
		Version:            Version,
	})
}

func CreateIfNotExists() stacktrace.Error {
	if _, err := os.Stat(File); os.IsNotExist(err) {
		return saveEmtpyConfig()
	}
	return nil
}

// Save replaces the config file via rename instead of truncating it in place. Every duckman invocation, including the
// shell completion sourced on each shell start, reads the config, so concurrent processes must never observe it empty
// or half-written.
func Save(config any) stacktrace.Error {
	serialized, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return stacktrace.Wrap(err)
	}

	tmp, err := os.CreateTemp(Dir, filepath.Base(File)+".*.tmp")
	if err != nil {
		return stacktrace.Wrap(err)
	}
	defer os.Remove(tmp.Name())

	if _, err := tmp.Write(serialized); err != nil {
		tmp.Close()
		return stacktrace.Wrap(err)
	}
	if err := tmp.Close(); err != nil {
		return stacktrace.Wrap(err)
	}
	return stacktrace.Wrap(os.Rename(tmp.Name(), File))
}
