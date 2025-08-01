package models

import "os"

type FileAppender struct {
	FilePath string
}

func NewFileAppender(filePath string) *FileAppender {
	return &FileAppender{
		FilePath: filePath,
	}
}

func (a *FileAppender) Append(message *LogMessage) error {
	f, err := os.OpenFile(a.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(message.String() + "\n")
	return err
}
