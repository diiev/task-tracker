package cli

import (
	"fmt"

	"go.mod/internal/model"
	"go.mod/internal/storage"
)

func Add(description string) error {

	if description == "" {
		return fmt.Errorf("Описание задачи не может быть пустым")
	}

	tasks, err := storage.LoadTasks("tasks.json")
	if err != nil {
		return fmt.Errorf("Ошибка загрузки файла")
	}

	newID := 1
	for _, t := range tasks {
		if t.ID >= newID {
			newID = t.ID + 1
		}
	}
	newTask := model.NewTask(newID, description)
	tasks = append(tasks, newTask)
	if err := storage.SaveTask(tasks); err != nil {
		return fmt.Errorf("Ошибка сохранения файла %w", err)
	}
	fmt.Println("Задача добавлена (id: %d): %s\n", newTask.ID, newTask.Description)
	return nil

}
