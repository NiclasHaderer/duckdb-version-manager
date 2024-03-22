package migrations

import (
	"github.com/hashicorp/go-version"
	"sort"
)

type UntypedConfig = map[string]any

type Migration struct {
	ToVersion version.Version
	Up        func(localConfig UntypedConfig) (UntypedConfig, error)
}

var migrations []Migration

func AddMigration(toVersion version.Version, up func(localConfig UntypedConfig) (UntypedConfig, error)) {
	migrations = append(migrations, Migration{ToVersion: toVersion, Up: up})
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].ToVersion.LessThan(&migrations[j].ToVersion)
	})
}
