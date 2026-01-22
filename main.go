package main

import (
	"bufio"
	"flag"
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

	/*
		// OLD LOGIC: Manual argument parsing
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
	*/

	// NEW LOGIC: Using 'flag' package for subcommands
	// and 'bufio' for interactive prompts

	// Define subcommands
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		//Args() returns the non-flag arguments
		if len(addCmd.Args()) == 0 {
			fmt.Println("Usage: task-cli add <description>")
			return
		}
		// Join remaining args to form the description
		description := strings.Join(addCmd.Args(), " ")
		err := task.AddTask(description)
		if err != nil {
			fmt.Printf("Error adding task: %v\n", err)
			return
		}
		fmt.Println("Task added successfully.")

	case "list":
		listCmd.Parse(os.Args[2:])
		err := task.ListTasks()
		if err != nil {
			fmt.Printf("Error listing tasks: %v\n", err)
		}

	case "complete":
		completeCmd.Parse(os.Args[2:])
		args := completeCmd.Args()
		if len(args) < 1 {
			fmt.Println("Usage: task-cli complete <id>")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID.")
			return
		}
		err = task.CompleteTask(id)
		if err != nil {
			fmt.Printf("Error completing task: %v\n", err)
			return
		}
		fmt.Println("Task marked as completed.")

	case "delete":
		deleteCmd.Parse(os.Args[2:])
		args := deleteCmd.Args()
		if len(args) < 1 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID.")
			return
		}

		// Interactive Confirmation using bufio
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Are you sure you want to delete task %d? (y/N): ", id)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))

		if response == "y" || response == "yes" {
			err = task.DeleteTask(id)
			if err != nil {
				fmt.Printf("Error deleting task: %v\n", err)
				return
			}
			fmt.Println("Task deleted successfully.")
		} else {
			fmt.Println("Deletion cancelled.")
		}

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
