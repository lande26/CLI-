package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"task-cli/task"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add <description>")
			return
		}
		description := strings.Join(os.Args[2:], " ")
		err := task.AddTask(description)
		if err != nil {
			fmt.Printf("Error adding task: %v\n", err)
			return
		}
		fmt.Println("Task added successfully.")

	case "list":
		err := task.ListTasks()
		if err != nil {
			fmt.Printf("Error listing tasks: %v\n", err)
		}

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli complete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID. It must be a number.")
			return
		}
		err = task.CompleteTask(id)
		if err != nil {
			fmt.Printf("Error completing task: %v\n", err)
			return
		}
		fmt.Println("Task marked as completed.")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID. It must be a number.")
			return
		}
		err = task.DeleteTask(id)
		if err != nil {
			fmt.Printf("Error deleting task: %v\n", err)
			return
		}
		fmt.Println("Task deleted successfully.")

	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Task Tracker CLI")
	fmt.Println("Usage:")
	fmt.Println("  add <description>   Add a new task")
	fmt.Println("  list                List all tasks")
	fmt.Println("  complete <id>       Mark a task as completed")
	fmt.Println("  delete <id>         Delete a task")
}
