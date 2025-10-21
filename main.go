package main

import (
	"fmt"
	"os"
	"strconv"

	"go.mod/internal/cli"
)

func main() {
	argTackTracker := os.Args
	if len(argTackTracker) > 1 {
		if argTackTracker[1] == "add" {
			cli.AddTask(argTackTracker[2])
		}
		if argTackTracker[1] == "list" {
			fmt.Println("Cписок задач:\n--------------------------")
			cli.ShowTaskList()
		}
		if argTackTracker[1] == "delete" {
			id, _ := strconv.Atoi(argTackTracker[2])
			cli.DeleteTask(id)
		}

	} else {
		fmt.Println("Аргументы не найдены")
	}

}
