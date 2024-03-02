package utils

import (
	"fmt"
	"os"
)

func ExitWith(format string, v ...any) {
	fmt.Printf(format, v...)
	fmt.Printf("\n")
	os.Exit(1)
}

func ExitWithError(err error) {
	errorType := fmt.Sprintf("%T", err)
	fmt.Printf("%s: %s\n", errorType, err)
	os.Exit(1)
}
