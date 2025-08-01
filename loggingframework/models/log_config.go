package models

type LoggerConfig struct {
	LogLevel    LogLevel
	LogAppender LogAppender
}

func NewLoggerConfig(level *LogLevel, logAppender *LogAppender) *LoggerConfig {
	return &LoggerConfig{
		LogLevel:    *level,
		LogAppender: *logAppender,
	}
}
