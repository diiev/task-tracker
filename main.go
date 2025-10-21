package main

import (
	"fmt"
	"os"

	"go.mod/internal/cli"
)

func main() {
	argTackTracker := os.Args
	if len(argTackTracker) > 1 {
		if argTackTracker[1] == "add" {
			fmt.Println("Add tasks")
		}
		if argTackTracker[1] == "list" {
			fmt.Println("Lists of tasks")
			cli.RunList()
		}

	} else {
		fmt.Println("Arguments not found")
	}

}
