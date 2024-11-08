package main

import (
	"fmt"
	version2 "github.com/hashicorp/go-version"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Call this util using only one argument")
	}

	version := os.Args[1]

	_, err := version2.NewVersion(version)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
