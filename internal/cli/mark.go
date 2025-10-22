package cli

import (
	"fmt"
	"time"

	"go.mod/internal/model"
	"go.mod/internal/storage"
)

func MarkTask(id int, status string) error {
	if status == "" {
		return fmt.Errorf("Статус должен быть: mark-in-progress или mark-done")
	}
	tasks, err := storage.LoadTasks("tasks.json")
	if err != nil {
		return fmt.Errorf("Ошибка загрузки данных %w", err)
	}
	found := false
	for _, t := range tasks {
		if t.ID == id {
			found = true
			if status == "mark-in-progress" {
				t.Status = string(model.TASK_STATUS_IN_PROGRESS)
			}
			if status == "mark-done" {
				t.Status = string(model.TASK_STATUS_DONE)
			}
			t.UpdatedAt = time.Now()
		}
	}
	if err := storage.SaveTask(tasks); err != nil {
		return fmt.Errorf("Ошибка сохранения файла %w", err)
	}
	if !found {
		fmt.Printf("Задача c id - %d не найдена\n", id)
	} else {
		fmt.Printf("Задача c id - %d помечена как %s\n", id, status[5:])
	}

	return nil

}
