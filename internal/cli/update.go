package cli

import (
	"fmt"
	"time"

	"go.mod/internal/storage"
)

func UpdateTask(id int, newDescription string) error {
	if newDescription == "" {
		return fmt.Errorf("Описание не должно быть пустым")
	}
	tasks, err := storage.LoadTasks("tasks.json")
	if err != nil {
		return fmt.Errorf("Ошибка загрузки данных %w", err)
	}

	found := false
	for _, t := range tasks {
		if t.ID == id {
			found = true
			t.Description = newDescription
			t.UpdatedAt = time.Now()
		}
	}
	if err := storage.SaveTask(tasks); err != nil {
		return fmt.Errorf("Ошибка сохранения файла %w", err)
	}
	if !found {
		fmt.Printf("Задача c id - %d не найдена\n", id)
	} else {
		fmt.Printf("Задача c id - %d обновлена\n", id)
	}
	return nil
}
