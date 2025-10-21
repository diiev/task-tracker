package cli

import (
	"fmt"

	"go.mod/internal/model"
	"go.mod/internal/storage"
)

func DeleteTask(id int) error {
	tasks, err := storage.LoadTasks("tasks.json")
	if err != nil {
		return fmt.Errorf("Ошибка загрузки данных %w", err)
	}
	found := false
	var deleteTasks []*model.Task
	for _, t := range tasks {
		if t.ID != id {
			deleteTasks = append(deleteTasks, t)
		} else {
			found = true
		}
	}
	if !found {
		return fmt.Errorf("Задача не найдена %w", err)
	}
	storage.SaveTask(deleteTasks)
	return nil
}
