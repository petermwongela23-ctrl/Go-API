package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func saveTasks(tasks []Task) {
	// json.MarshalIndent formats the JSON neatly with indentation
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	// Write the data to a file with read/write permissions (0644)
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

// loadTasks reads tasks.json from disk when the program starts
func loadTasks() []Task {
	// Read the file if it exists
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		// If the file doesn't exist yet, just return an empty slice
		return []Task{}
	}

	var tasks []Task
	// Unpack (deserialize) the JSON data back into our Task struct slice
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error reading saved tasks:", err)
		return []Task{}
	}

	return tasks
}

// Task represents a single task item
type Task struct {
	ID          int
	Description string
	Completed   bool
}

// addTask creates a new task and appends it to our slice
func addTask(tasks []Task, description string) []Task {
	newID := len(tasks) + 1

	newTask := Task{
		ID:          newID,
		Description: description,
		Completed:   false,
	}

	return append(tasks, newTask)
}

// listTasks displays all current tasks in the terminal
func listTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("\nNo tasks found.")
		return
	}

	fmt.Println("\nYour Tasks:")
	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "x"
		}
		fmt.Printf("[%s] %d. %s\n", status, task.ID, task.Description)
	}
	fmt.Println()
}

// completeTask finds a task by its ID and marks it as completed
func completeTask(tasks []Task, id int) []Task {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			fmt.Printf("\nTask %d marked as completed!\n", id)
			return tasks
		}
	}
	fmt.Printf("\nTask with ID %d not found.\n", id)
	return tasks
}

// deleteTask removes a task from our slice by its ID
func deleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			// Go trick to remove an item from a slice using append and slicing
			updatedTasks := append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("\nTask %d deleted successfully!\n", id)
			return updatedTasks
		}
	}
	fmt.Printf("\nTask with ID %d not found.\n", id)
	return tasks
}

func main() {
	var tasks []Task
	reader := bufio.NewReader(os.Stdin)

	for {
		// Display Menu Options
		fmt.Println("--- Task Manager ---")
		fmt.Println("1. List Tasks")
		fmt.Println("2. Add Task")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		// Read user input choice
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			listTasks(tasks)

		case "2":
			fmt.Print("Enter task description: ")
			desc, _ := reader.ReadString('\n')
			desc = strings.TrimSpace(desc)
			if desc != "" {
				tasks = addTask(tasks, desc)
				fmt.Println("Task added successfully!")
			}

		case "3":
			listTasks(tasks)
			fmt.Print("Enter task ID to mark complete: ")
			idInput, _ := reader.ReadString('\n')
			idInput = strings.TrimSpace(idInput)
			idNum, err := strconv.Atoi(idInput)
			if err == nil {
				tasks = completeTask(tasks, idNum)
			} else {
				fmt.Println("Invalid ID. Please enter a number.")
			}
		case "4": // New case for deletion
			listTasks(tasks)
			fmt.Print("Enter task ID to delete: ")
			idInput, _ := reader.ReadString('\n')
			idInput = strings.TrimSpace(idInput)
			idNum, err := strconv.Atoi(idInput)
			if err == nil {
				tasks = deleteTask(tasks, idNum)
			} else {
				fmt.Println("Invalid ID. Please enter a number.")
			}
		case "5":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice, please choose between 1 and 4.")
		}
		fmt.Println()
	}
}
