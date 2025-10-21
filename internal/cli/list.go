package cli

import (
	"fmt"
	"log"

	"go.mod/internal/storage"
)

func RunList() {
	tasks, err := storage.LoadTasks("tasks.json")
	if err != nil {
		log.Fatal(err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	for _, t := range tasks {
		if t != nil {
			fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreatedAt: %s\nUpdatedAt: %s\n\n",
				t.ID, t.Description, t.Status, t.CreatedAt.Format("2006-01-02 15:04"), t.UpdatedAt.Format("2006-01-02 15:04"))
		}
	}
}
