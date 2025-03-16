package main

import (
	"fmt"
	"os"
	"whoami/platform"
)

func printUsage() {
	fmt.Println("usage: whoami")
}

func main() {
	if len(os.Args) > 1 {
		printUsage()
		os.Exit(1)
	}

	user, err := platform.GetUserId()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(user)
}
