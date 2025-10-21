package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"go.mod/internal/model"
)

// LoadTasks — загружает задачи из файла tasks.json или создаёт новый, если его нет
func LoadTasks(filename string) ([]*model.Task, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {

		empty := []*model.Task{}
		data, err := json.MarshalIndent(empty, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("ошибка сериализации: %w", err)
		}

		if err := os.WriteFile(filename, data, 0644); err != nil {
			return nil, fmt.Errorf("ошибка создания файла: %w", err)
		}

		return empty, nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	if len(data) == 0 {
		return []*model.Task{}, nil
	}

	var tasks []*model.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %w", err)
	}

	return tasks, nil
}
