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
			cli.Add(argTackTracker[2])
		}
		if argTackTracker[1] == "list" {
			fmt.Println("Cписок задач:\n--------------------------")
			cli.RunList()
		}

	} else {
		fmt.Println("Аргументы не найдены")
	}

}
