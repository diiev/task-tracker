package cli

import (
	"fmt"

	"go.mod/internal/model"
	"go.mod/internal/storage"
)

func DeleteTask(id int) error {
	tasks, err := storage.LoadTasks(storage.TaskFile)
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
	if err := storage.SaveTask(deleteTasks); err != nil {
		return fmt.Errorf("Ошибка сохранения файла %w", err)
	}
	return nil
}
