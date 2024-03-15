package utils

import (
	"duckdb-version-manager/stacktrace"
	"os"
)

func RemoveFileOrDie(file string) {
	if _, err := os.Stat(file); err == nil {
		err := os.RemoveAll(file)
		if err != nil {
			ExitWithError(stacktrace.Wrap(err))
		}
	}
}
