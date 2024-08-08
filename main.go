package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Title string
	Done  bool
}

func main() {
	tasks := []Task{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("To-Do List")
		fmt.Println("1. Add a new task")
		fmt.Println("2. View tasks")
		fmt.Println("3. Mark a task as done")
		fmt.Println("4. Delete a task")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addTask(&tasks, reader)
		case "2":
			viewTasks(tasks)
		case "3":
			markTaskDone(&tasks, reader)
		case "4":
			deleteTask(&tasks, reader)
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
func addTask(tasks *[]Task, reader *bufio.Reader) {
	fmt.Print("Enter task title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	*tasks = append(*tasks, Task{Title: title, Done: false})
	fmt.Println("Task added!")
}
func viewTasks(tasks []Task) {
	fmt.Println("Tasks:")
	for i, task := range tasks {
		status := "Not Done"
		if task.Done {
			status = "Done"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.Title, status)
	}
}
func markTaskDone(tasks *[]Task, reader *bufio.Reader) {
	viewTasks(*tasks)
	fmt.Print("Enter task number to mark as done: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	index := parseIndex(input, len(*tasks))
	if index >= 0 {
		(*tasks)[index].Done = true
		fmt.Println("Task marked as done!")
	} else {
		fmt.Println("Invalid task number.")
	}
}
func deleteTask(tasks *[]Task, reader *bufio.Reader) {
	viewTasks(*tasks)
	fmt.Print("Enter task number to delete: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	index := parseIndex(input, len(*tasks))
	if index >= 0 {
		*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
		fmt.Println("Task deleted!")
	} else {
		fmt.Println("Invalid task number.")
	}
}
func parseIndex(input string, length int) int {
	index := -1
	fmt.Sscanf(input, "%d", &index)
	index -= 1 // Convert to zero-based index
	if index < 0 || index >= length {
		return -1
	}
	return index
}
