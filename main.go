package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type TaskList struct {
	ID   int
	Task Task
}

type Task struct {
	Description string
	Status      string
	createdAt   time.Time
	updatedAt   time.Time
}

var taskList []TaskList

func main() {

	fmt.Println("Welcome to Task CLI App!")
	fmt.Println("What would you like to do?")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "q" {
			return
		}
		withQuotes := strings.Split(scanner.Text(), " \"")
		operations := strings.Split(withQuotes[0], " ")

		fmt.Println("Operations:", operations[0])
		// if len(command) <= 1 {
		// 	fmt.Println("You've entered incorrect commands (e.g. add groceries)")
		// 	continue
		// }
		// fmt.Println("Operation:", command[0])
		// fmt.Println("Description:", strings.Trim(command[1], "\""))
		// fmt.Println("You wrote:", scanner.Text())

		// switch command[0] {
		// case "add":
		// 	fmt.Println("add called")
		// 	addTask(strings.Trim(command[1], "\""))
		// case "update":
		// 	fmt.Println("update called")
		// case "delete":
		// 	fmt.Println("delete called")
		// case "list":
		// 	fmt.Println("list called")
		// 	listTasks()
		// default:
		// 	fmt.Println("Incorrect operations")
		// 	continue
		// }

	}
}

func addTask(description string) {
	trimmedDesc := strings.Trim(description, "\"")

	task := Task{
		Description: trimmedDesc,
		Status:      "Pending",
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}

	fmt.Println(task)

	tL := TaskList{
		ID:   1,
		Task: task,
	}

	taskList = append(taskList, tL)

}

func listTasks() {
	f, err := os.Open("tasklist.json")

	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	decoder.Token()

	data := map[string]any{}
	for decoder.More() {
		decoder.Decode(&data)

	}

	fmt.Println(data)
}
