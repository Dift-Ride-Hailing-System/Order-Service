package logger

import (
	"fmt"
	"log"
	"time"
)

func formatMessage(level, msg string, args ...interface{}) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	return fmt.Sprintf("[%s] [%s] %s", timestamp, level, msg)
}

// Info log
func Info(msg string, args ...interface{}) {
	log.Println(formatMessage("INFO", msg, args...))
}

// Warn log
func Warn(msg string, args ...interface{}) {
	log.Println(formatMessage("WARN", msg, args...))
}

// Error log
func Error(msg string, args ...interface{}) {
	log.Println(formatMessage("ERROR", msg, args...))
}

// Debug log (optionally can be filtered in production)
func Debug(msg string, args ...interface{}) {
	log.Println(formatMessage("DEBUG", msg, args...))
}
