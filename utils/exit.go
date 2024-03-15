package utils

import (
	"duckdb-version-manager/stacktrace"
	"fmt"
	"os"
)

func exitWith(format string, v ...any) {
	fmt.Printf(format, v...)
	fmt.Printf("\n")
	os.Exit(1)
}

func ExitWithError(err stacktrace.Error) {
	var errorMsg string

	// Check if DUCKMAN_DEBUG is set
	if EnvIsTruthy("DUCKMAN_DEBUG") {
		errorMsg = fmt.Sprintf("%s\n%s", err, err.StackTrace())
	} else {
		errorMsg = fmt.Sprintf("%s", err)
	}

	exitWith(errorMsg)
}
