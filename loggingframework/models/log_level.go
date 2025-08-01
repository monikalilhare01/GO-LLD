package models

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
	LogLevelFatal
)

func (l LogLevel) string() string {
	return [...]string{"DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}[l]
}
