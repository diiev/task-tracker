package cli

import (
	"fmt"
	"log"

	"go.mod/internal/model"
	"go.mod/internal/storage"
)

func ShowTaskList(status string) {
	tasks, err := storage.LoadTasks("tasks.json")
	if err != nil {
		log.Fatal(err)
	}
	var filtered []*model.Task

	for _, t := range tasks {
		if status == "done" && t.Status == string(model.TASK_STATUS_DONE) {
			filtered = append(filtered, t)
		}
		if status == "" {
			filtered = append(filtered, t)
		}
		if status == "in-progress" && t.Status == string(model.TASK_STATUS_IN_PROGRESS) {
			filtered = append(filtered, t)
		}

	}

	for _, t := range filtered {
		if t != nil {
			fmt.Println("-------------------------------")
			fmt.Printf("Номер задачи: %d\nОписание: %s\nСтатус: %s\nСоздана: %s\nОбновлено: %s\n",
				t.ID, t.Description, t.Status, t.CreatedAt.Format("2006-01-02 15:04"), t.UpdatedAt.Format("2006-01-02 15:04"))
		} else {
			fmt.Println("Задач нет")
		}
	}
	if len(filtered) == 0 {
		fmt.Println("Задачи не найдены")
		return
	}
}
