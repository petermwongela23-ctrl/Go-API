package main

import "fmt"

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
		fmt.Println("No tasks found.")
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
func completeTask(tasks []Task, id int) []Task {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			fmt.Printf("Task %d marked as completed!\n", id)
			return tasks
		}
	}
	fmt.Printf("Task with ID %d not found.\n", id)
	return tasks
}
func main() {
	var tasks []Task

	tasks = addTask(tasks, "Learn Go fundamentals")
	tasks = addTask(tasks, "Build a CLI task manager")
	tasks = completeTask(tasks, 1)
	// Display the tasks
	listTasks(tasks)
}
