package models

import "fmt"

type ConsoleAppender struct{}

func NewConsoleAppender() *ConsoleAppender {
	return &ConsoleAppender{}
}

func (ca *ConsoleAppender) Append(msg *LogMessage) error {
	fmt.Println(msg.String())
	return nil
}
