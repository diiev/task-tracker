package cmd

import (
	"fmt"
	"os"
	"strconv"

	"go.mod/internal/cli"
)

func Run() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Недостаточно аргументов для команды add.\nИспользование: task-tracker add \"описание задачи\"")
			return
		}
		description := os.Args[2]
		cli.AddTask(description)

	case "list":
		status := ""
		if len(os.Args) > 2 {
			status = os.Args[2]
		}
		fmt.Println("Список задач:\n")
		cli.ShowTaskList(status)

	case "delete", "update", "mark-in-progress", "mark-done":
		if len(os.Args) < 3 {
			fmt.Printf("Недостаточно аргументов для команды %s\n", command)
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID должен быть числом.")
			return
		}

		switch command {
		case "delete":
			cli.DeleteTask(id)
		case "update":
			if len(os.Args) < 4 {
				fmt.Println("Недостаточно аргументов для команды update.\nИспользование: task-tracker update <id> \"новое описание\"")
				return
			}
			newDescription := os.Args[3]
			cli.UpdateTask(id, newDescription)
		case "mark-in-progress", "mark-done":
			cli.MarkTask(id, command)
		}

	default:
		showUsage()
	}
}

func showUsage() {
	fmt.Println(`🧭 Task Tracker CLI (Go)
Простое CLI-приложение для управления задачами на Go.

Использование:
  task-tracker <command> [arguments]

Команды:
  add "описание"              Добавить новую задачу
  list [status]               Показать задачи (все или по статусу)
  update <id> "описание"      Обновить описание задачи
  delete <id>                 Удалить задачу
  mark-in-progress <id>       Отметить задачу как "в процессе"
  mark-done <id>              Отметить задачу как "выполнена"

Примеры:
  task-tracker add "Купить хлеб"
  task-tracker list
  task-tracker list done
  task-tracker update 1 "Купить хлеб и молоко"
  task-tracker mark-in-progress 1
  task-tracker mark-done 1
`)
}
