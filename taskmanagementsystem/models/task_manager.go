package models

import "sync"

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

func (tm *TaskManger) assignTaskToUser(user *User, task *Task) {
	tm.userTask[user.GetID()] = append(tm.userTask[user.GetID()], task)
}
