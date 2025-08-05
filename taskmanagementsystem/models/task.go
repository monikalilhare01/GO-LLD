package models

import "time"

type Task struct {
	id           string
	title        string
	description  string
	dueDate      time.Time
	assignedUser *User
	priority     TaskPriority
	status       TaskStatus
}

func NewTask(id string, title, description string, dueDate time.Time, assignedUser *User, priority TaskPriority, status TaskStatus) *Task {
	return &Task{
		id:           id,
		title:        title,
		description:  description,
		dueDate:      dueDate,
		assignedUser: assignedUser,
		priority:     priority,
		status:       status,
	}
}

func (t *Task) GetID() string {
	return t.id
}

func (t *Task) GetTitle() string {
	return t.title
}

func (t *Task) GetDescription() string {
	return t.description
}

func (t *Task) GetDueDate() time.Time {
	return t.dueDate
}

func (t *Task) GetAssignedUser() *User {
	return t.assignedUser
}

func (t *Task) GetPriority() TaskPriority {
	return t.priority
}

func (t *Task) GetStatu() TaskStatus {
	return t.status
}

func (t *Task) SetID(id string) {
	t.id = id
}

func (t *Task) SetTitle(title string) {
	t.title = title
}

func (t *Task) SetDescription(desc string) {
	t.description = desc
}

func (t *Task) SetDueDate(date time.Time) {
	t.dueDate = date
}

func (t *Task) SetPriority(priority TaskPriority) {
	t.priority = priority
}

func (t *Task) SetStatus(status TaskStatus) {
	t.status = status
}
