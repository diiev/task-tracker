package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"go.mod/internal/model"
)

const TaskFile = "tasks.json"

// SaveTask - сохраняет задачи
func SaveTask(tasks []*model.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return fmt.Errorf("Ошибка записи JSON файла %w", err)
	}
	tmpFile := TaskFile + ".tmp"

	f, err := os.Create(tmpFile)
	if err != nil {
		return fmt.Errorf("Ошибка создания временного файла %w", err)
	}
	_, err = f.Write(data)
	if err != nil {
		f.Close()
		return fmt.Errorf("Ошибка записи файла %w", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("Ошибка закртия временного файла")
	}
	if err := os.Rename(tmpFile, TaskFile); err != nil {
		return fmt.Errorf("Ошибка переименования файла %w", err)
	}
	defer func() {
		f.Close()
	}()
	return nil
}

// LoadTasks — загружает задачи из файла  или создаёт новый, если его нет
func LoadTasks(filename string) ([]*model.Task, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {

		empty := []*model.Task{}
		data, err := json.MarshalIndent(empty, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("Ошибка сериализации: %w", err)
		}

		if err := os.WriteFile(filename, data, 0644); err != nil {
			return nil, fmt.Errorf("Ошибка создания файла: %w", err)
		}

		return empty, nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения файла: %w", err)
	}

	if len(data) == 0 {
		return []*model.Task{}, nil
	}

	var tasks []*model.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("Ошибка парсинга JSON: %w", err)
	}

	return tasks, nil
}
