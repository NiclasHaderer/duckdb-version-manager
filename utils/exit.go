package utils

import (
	"fmt"
	"os"
)

func ExitWith(format string, v ...any) {
	fmt.Printf("Duck-VM encounted a fatal error\n")
	fmt.Printf(format, v...)
	fmt.Printf("\n")
	os.Exit(1)
}

func ExitWithError(err error) {
	errorType := fmt.Sprintf("%T", err)
	errorMsg := fmt.Sprintf("%s: %s", errorType, err)
	ExitWith(errorMsg)
}
