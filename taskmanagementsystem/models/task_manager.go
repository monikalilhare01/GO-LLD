package models

import (
	"sync"
	"time"
)

type TaskManger struct {
	task     map[string]*Task
	userTask map[string][]*Task
	mu       sync.Mutex
}

var instance *TaskManger
var once sync.Once

func GetTaskManager() *TaskManger {
	once.Do(func() {
		instance = &TaskManger{
			task:     make(map[string]*Task),
			userTask: make(map[string][]*Task),
		}
	})
	return instance
}

func (tm *TaskManger) CreateTask(t *Task) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	tm.task[t.GetID()] = t
	tm.assignTaskToUser(t.GetAssignedUser(), t)
}

func (tm *TaskManger) UpdateTask(t *Task) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if fetchedTask, exist := tm.task[t.GetID()]; exist {
		fetchedTask.SetTitle(t.GetTitle())
		fetchedTask.SetDescription(t.GetDescription())
		fetchedTask.SetDueDate(t.GetDueDate())
		fetchedTask.SetPriority(t.GetPriority())
		fetchedTask.SetStatus(t.GetStatu())

		if fetchedTask.GetAssignedUser() != t.GetAssignedUser() {
			tm.unassignTaskFromUser(fetchedTask.assignedUser, fetchedTask)
			tm.assignTaskToUser(fetchedTask.GetAssignedUser(), fetchedTask)
		}
	}
}

func (tm *TaskManger) DeleteTask(taskId string) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if task, exits := tm.task[taskId]; exits {
		delete(tm.task, taskId)
		tm.unassignTaskFromUser(task.GetAssignedUser(), task)
	}
}

func (tm *TaskManger) SearchTask(keyword string) []*Task {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	var MatchingTask []*Task
	for _, task := range tm.task {
		if contains(task.GetDescription(), keyword) || contains(task.GetTitle(), keyword) {
			MatchingTask = append(MatchingTask, task)
		}
	}
	return MatchingTask
}

func (tm *TaskManger) FilterTask(status TaskStatus, startDate, endDate time.Time, priority int) []*Task {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	var filteredTask []*Task

	for _, task := range tm.task {
		if task.GetStatu() == status &&
			task.GetDueDate().After(startDate) &&
			task.GetDueDate().Before(endDate) &&
			task.GetPriority() == task.priority {
			filteredTask = append(filteredTask, task)
		}
	}

	return filteredTask
}

func (tm *TaskManger) MarkTaskAsCompleted(taskId string) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if task, exists := tm.task[taskId]; exists {
		task.SetStatus(DONE)
	}
}

func (tm *TaskManger) GetTaskHistory(user *User) []*Task {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	return tm.userTask[user.GetID()]
}

func contains(text, substr string) bool {
	return len(text) >= len(substr) && text[:len(substr)] == substr
}

func (tm *TaskManger) unassignTaskFromUser(user *User, task *Task) {
	tasks := tm.userTask[user.GetID()]
	for i, t := range tasks {
		if t == task {
			tm.userTask[user.GetID()] = append(tasks[:i], tasks[i+1:]...)
		}
	}
}

func (tm *TaskManger) assignTaskToUser(user *User, task *Task) {
	tm.userTask[user.GetID()] = append(tm.userTask[user.GetID()], task)
}
