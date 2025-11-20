package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type TaskList struct {
	ID   uuid.UUID `json:"id"`
	Task Task      `json:"task"`
}

type Task struct {
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var taskList []TaskList

func main() {

	fmt.Println("Welcome to Task CLI App!")
	fmt.Print("\n List of Operations ----------------\n")
	fmt.Print("\n1. add")
	fmt.Print("\n2. delete")
	fmt.Print("\n3. list")
	fmt.Print("\n\nWhat would you like to do? Press 'q' to quit\n")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "q" {
			return
		}

		userInput := strings.ToLower(scanner.Text())

		// if len(command) <= 1 {
		// 	fmt.Println("You've entered incorrect commands (e.g. add groceries)")
		// 	continue
		// }
		switch userInput {
		case "add":
			addTask()
		case "update":
			fmt.Println("update called")
		case "delete":
			deleteTask()
		case "list":
			listTasks()
		default:
			fmt.Println("Incorrect operations")
			continue
		}

		fmt.Print("\n\nWhat would you like to do? Press 'q' to quit\n")
	}
}

func deleteTask() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What would you like to delete?")
	if scanner.Scan() {
		userInput := scanner.Text()

		taskList := []TaskList{}

		fileContent, err := os.ReadFile("tasklist.json")
		if err != nil {
			fmt.Println(err)
		}

		if len(fileContent) > 0 {
			err = json.Unmarshal(fileContent, &taskList)
			if err != nil {
				fmt.Println(err)
			}
		}
		targetID, err := uuid.Parse(userInput)
		if err != nil {
			fmt.Println(err)
		}
		updatedTaskList := []TaskList{}
		for _, task := range taskList {
			if (task.ID) == targetID {
				continue
			}
			updatedTaskList = append(updatedTaskList, task)
		}
		jsonData, err := json.MarshalIndent(updatedTaskList, "", " ")
		if err != nil {
			fmt.Println(err)
		}
		err = os.WriteFile("tasklist.json", jsonData, 6044)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func addTask() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("what would you like to add?")

	if scanner.Scan() {
		userInput := scanner.Text()

		task := Task{
			Description: userInput,
			Status:      "Pending",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		tL := TaskList{
			ID:   uuid.New(),
			Task: task,
		}
		taskList = []TaskList{}
		fileContent, err := os.ReadFile("tasklist.json")
		if err != nil {
			fmt.Println(err)
		}

		if len(fileContent) > 0 {
			err = json.Unmarshal(fileContent, &taskList)
			if err != nil {
				fmt.Println(err)
			}
		}
		taskList = append(taskList, tL)

		jsonData, err := json.MarshalIndent(taskList, "", " ")

		if err != nil {
			fmt.Println(err)
		}

		err = os.WriteFile("tasklist.json", jsonData, 0644)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Print("\n")
		fmt.Println("Task has been added")
		fmt.Println("ID:", tL.ID)
		fmt.Println("Description:", tL.Task.Description)

	}

}

func listTasks() {
	filePath := "tasklist.json"

	tasks := []TaskList{}
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	if len(fileContent) > 0 {
		err = json.Unmarshal(fileContent, &tasks)
		if err != nil {
			fmt.Println(err)
		}
	}

	for _, task := range tasks {
		fmt.Print("\n")
		fmt.Println(task.ID)
		fmt.Println("Description:", task.Task.Description)
		fmt.Println("Status:", task.Task.Status)
	}
}
