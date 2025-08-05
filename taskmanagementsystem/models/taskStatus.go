package models

type TaskStatus string

const (
	TODO       TaskStatus = "TODO"
	INPROGRESS TaskStatus = "In-Progress"
	DONE       TaskStatus = "Done"
	BLOCKED    TaskStatus = "Blocked"
)
