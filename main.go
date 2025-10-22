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
			status := ""
			fmt.Println("Cписок задач:\n")
			if len(argTackTracker) > 2 {
				status = argTackTracker[2]
			}
			cli.ShowTaskList(status)
		}
		if argTackTracker[1] == "delete" {
			id, _ := strconv.Atoi(argTackTracker[2])
			cli.DeleteTask(id)
		}
		if argTackTracker[1] == "update" {
			id, _ := strconv.Atoi(argTackTracker[2])
			cli.UpdateTask(id, argTackTracker[3])
		}
		if argTackTracker[1] == "mark-in-progress" {
			id, _ := strconv.Atoi(argTackTracker[2])
			cli.MarkTask(id, argTackTracker[1])
		}
		if argTackTracker[1] == "mark-done" {
			id, _ := strconv.Atoi(argTackTracker[2])
			cli.MarkTask(id, argTackTracker[1])
		}

	} else {
		fmt.Println("Аргументы не найдены")
	}

}
