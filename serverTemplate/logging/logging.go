package logging

import (
	"log"
	"os"
)

var (
	enableDebug = false
	infoLog     *log.Logger
	debugLog    *log.Logger
	errorLog    *log.Logger
	requestLog  *log.Logger
)

func init() {
	if debugFlag := os.Getenv("ENABLE_DEBUG"); debugFlag != "" {
		enableDebug = true
	}
	infoLog = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	debugLog = log.New(os.Stdout, "DEBUG: ", log.LstdFlags)
	errorLog = log.New(os.Stderr, "ERROR: ", log.LstdFlags)
	requestLog = log.New(os.Stdout, "", 0)
}

// Infof logs info level logs.
func Infof(format string, args ...interface{}) {
	infoLog.Printf(format, args...)
}

// Errorf logs error messages.
func Errorf(format string, args ...interface{}) {
	errorLog.Printf(format, args...)
}

// Debugf only logs in development
func Debugf(format string, args ...interface{}) {
	if enableDebug {
		debugLog.Printf(format, args...)
	}
}

// Requestf is for request logs.
func Requestf(format string, args ...interface{}) {
	requestLog.Printf(format, args...)
}
