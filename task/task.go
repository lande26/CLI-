package task

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
	Done        bool      `json:"done"`
}

const fileName = "tasks.json"

func LoadTasks() ([]Task, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}

func AddTask(description string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	newTask := Task{
		ID:          id,
		Description: description,
		CreatedAt:   time.Now(),
		Done:        false,
	}

	tasks = append(tasks, newTask)
	return SaveTasks(tasks)
}

func ListTasks() error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	fmt.Println("ID  Done  Description")
	for _, t := range tasks {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		fmt.Printf("%-3d %s   %s\n", t.ID, status, t.Description)
	}
	return nil
}

func CompleteTask(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Done = true
			tasks[i].CompletedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}

	return SaveTasks(tasks)
}

func DeleteTask(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	newTasks := []Task{}
	found := false
	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue
		}
		newTasks = append(newTasks, t)
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}

	return SaveTasks(newTasks)
}
