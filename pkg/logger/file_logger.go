package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var fileLogger *FileLogger
var once sync.Once

type FileLogger struct {
	AppName string
	Debug   bool
	logger  *log.Logger
	logFile *os.File
}

func UseBasicLogger() *FileLogger {
	if fileLogger == nil {
		log.Fatal("FileLogger is not initialized")
	}
	return fileLogger
}

func InitFileLogger(appName string, filename string, debug bool) *FileLogger {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error during initialization of logfile %s", err.Error())
	}

	logger := log.New(f, "", log.LstdFlags)
	logger.SetFlags(log.LstdFlags)

	once.Do(func() {
		fileLogger = &FileLogger{
			AppName: appName,
			logger:  logger,
			logFile: f,
			Debug:   debug,
		}
	})

	return fileLogger
}

func (l *FileLogger) Shutdown() {
	l.logFile.Close()
}

func (l *FileLogger) Info(message string, ctx LogContext) {
	logMessage := fmt.Sprintf("[INFO] %s; CONTEXT: %v", message, ctx)
	l.logger.Println(logMessage)

	if l.Debug {
		fmt.Fprintln(os.Stdout, logMessage)
	}
}

func (l *FileLogger) Error(message string, err error, ctx LogContext) {
	logMessage := fmt.Sprintf("[ERROR]: %s; GO_MESSAGE: %s; CONTEXT: %v", message, err.Error(), ctx)
	l.logger.Println(logMessage)

	if l.Debug {
		fmt.Fprintln(os.Stderr, logMessage)
	}
}

func (l *FileLogger) Warning(message string, ctx LogContext) {
	logMessage := fmt.Sprintf("[WARNING] %s; CONTEXT: %v", message, ctx)
	l.logger.Println(logMessage)

	if l.Debug {
		fmt.Fprintln(os.Stdout, logMessage)
	}
}
