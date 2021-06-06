package utils

import (
	"log"
)

const (
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

func Info(message string) {
	Log(INFO, message)
}

func Warn(message string) {
	Log(WARN, message)
}

func Err(message string) {
	Log(ERROR, message)
}

func Log(level string, message string) {
	log.Printf("[%s]: %s", level, message)
}
