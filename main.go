package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var tasks []Task
var filename = "tasks.json"

func loadTasks() {
	file, err := ioutil.ReadFile(filename)
	if err == nil {
		_ = json.Unmarshal(file, &tasks)
	}
}

func saveTasks() {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	_ = ioutil.WriteFile(filename, data, 0644)
}

func addTask(name string) {
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}
	task := Task{ID: id, Name: name, Done: false}
	tasks = append(tasks, task)
	saveTasks()
	fmt.Println("Task added:", name)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	fmt.Println("Tasks:")
	for _, task := range tasks {
		status := "Pending"
		if task.Done {
			status = "Done"
		}
		fmt.Printf("[%d] %s - %s\n", task.ID, task.Name, status)
	}
}

func markTaskDone(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			saveTasks()
			fmt.Println("Task marked as done:", task.Name)
			return
		}
	}
	fmt.Println("Task not found with ID:", id)
}

func deleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks()
			fmt.Println("Task deleted:", task.Name)
			return
		}
	}
	fmt.Println("Task not found with ID:", id)
}

func displayHelp() {
	fmt.Println("To-Do List CLI Application")
	fmt.Println("Commands:")
	fmt.Println("  add <task name>      - Add a new task")
	fmt.Println("  list                 - List all tasks")
	fmt.Println("  done <task ID>       - Mark a task as done")
	fmt.Println("  delete <task ID>     - Delete a task")
	fmt.Println("  help                 - Show this help message")
	fmt.Println("  exit                 - Exit the application")
}

func main() {
	loadTasks()
	displayHelp()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nEnter command: ")

		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		command := parts[0]
		args := strings.Join(parts[1:], " ")

		switch command {
		case "add":
			if args == "" {
				fmt.Println("Please provide a task name.")
			} else {
				addTask(args)
			}
		case "list":
			listTasks()
		case "done":
			if id, err := strconv.Atoi(args); err == nil {
				markTaskDone(id)
			} else {
				fmt.Println("Invalid task ID")
			}
		case "delete":
			if id, err := strconv.Atoi(args); err == nil {
				deleteTask(id)
			} else {
				fmt.Println("Invalid task ID")
			}
		case "help":
			displayHelp()
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid command. Type 'help' for a list of commands.")
		}
	}
}

