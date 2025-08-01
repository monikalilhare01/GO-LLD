package models

type LogAppender interface {
	Append(message *LogMessage) error
}
