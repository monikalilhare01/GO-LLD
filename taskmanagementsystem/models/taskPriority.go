package models

type TaskPriority int

const (
	LOW TaskPriority = iota
	MEDIUM
	HIGH
	CRITICAL
)
