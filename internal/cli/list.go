package cli

import (
	"fmt"
	"log"

	"go.mod/internal/storage"
)

func ShowTaskList() {
	tasks, err := storage.LoadTasks("tasks.json")
	if err != nil {
		log.Fatal(err)
	}

	if len(tasks) == 0 {
		fmt.Println("Задачи не найдены")
		return
	}

	for _, t := range tasks {
		if t != nil {
			fmt.Printf("Номер задачи: %d\nОписание: %s\nСтатус: %s\nСоздана: %s\nОбновлено: %s\n",
				t.ID, t.Description, t.Status, t.CreatedAt.Format("2006-01-02 15:04"), t.UpdatedAt.Format("2006-01-02 15:04"))
		}
		fmt.Println("-------------------------------")
	}
}
