package main

import (
	"fmt"
	"time"

	"github.com/monikalilhare01/GO-LLD/taskmanagementsystem/models"
)

func main() {
	taskManager := models.GetTaskManager()

	user1 := models.NewUser("1", "John Doe", "john@example.com")
	user2 := models.NewUser("2", "Jane Smith", "jane@example.com")

	task1 := models.NewTask("1", "Task 1", "Description 1", time.Now(), user1, models.MEDIUM, models.TODO)
	task2 := models.NewTask("2", "Task 2", "Description 2", time.Now(), user2, models.MEDIUM, models.INPROGRESS)
	task3 := models.NewTask("3", "Task 3", "Description 3", time.Now(), user1, models.MEDIUM, models.TODO)

	taskManager.CreateTask(task1)
	taskManager.CreateTask(task2)
	taskManager.CreateTask(task3)

	task2.SetDescription("Updated description")
	taskManager.UpdateTask(task2)

	searchResults := taskManager.SearchTask("Task")
	fmt.Println("Search Results:")
	for _, task := range searchResults {
		fmt.Println(task.GetTitle())
	}

	filteredTasks := taskManager.FilterTask(models.INPROGRESS, time.Unix(0, 0), time.Now(), 1)
	fmt.Println("Filtered Tasks:")
	for _, task := range filteredTasks {
		fmt.Println(task.GetTitle())
	}

	taskManager.MarkTaskAsCompleted("1")

	// Get task history for a user
	taskHistory := taskManager.GetTaskHistory(user1)
	fmt.Printf("Task History for %s:\n", user1.GetName())
	for _, task := range taskHistory {
		fmt.Println(task.GetTitle())
	}

	// Delete a task
	taskManager.DeleteTask("3")
}
