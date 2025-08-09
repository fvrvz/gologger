package gologger

import (
	"fmt"
	"log"
	"os"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
)

var (
	baseLogger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
)

func logWithColor(color, level, msg string, args ...interface{}) {
	prefix := fmt.Sprintf("%s[%s]%s ", color, level, colorReset)
	fullMsg := fmt.Sprintf(msg, args...)
	baseLogger.Output(3, prefix+fullMsg)
}

func INFO(msg string, args ...interface{}) {
	logWithColor(colorGreen, "INFO", msg, args...)
}

func DEBUG(msg string, args ...interface{}) {
	logWithColor(colorCyan, "DEBUG", msg, args...)
}

func WARN(msg string, args ...interface{}) {
	logWithColor(colorYellow, "WARN", msg, args...)
}

func ERROR(msg string, args ...interface{}) {
	logWithColor(colorRed, "ERROR", msg, args...)
}
