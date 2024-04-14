package utils

import (
	"duckdb-version-manager/stacktrace"
	"fmt"
	"os"
)

var beforeExitHooks []func(err stacktrace.Error)

func BeforeErrorExit(f func(err stacktrace.Error)) {
	beforeExitHooks = append(beforeExitHooks, f)
}

func exitWith(format string, err stacktrace.Error) {
	fmt.Print(format)
	fmt.Print("\n")

	for _, f := range beforeExitHooks {
		f(err)
	}
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

	exitWith(errorMsg, err)
}
