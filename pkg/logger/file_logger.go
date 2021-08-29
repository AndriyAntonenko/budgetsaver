package logger

import (
	"log"
	"os"
	"sync"
)

var fileLogger *FileLogger
var once sync.Once

type FileLogger struct {
	AppName string
	logger  *log.Logger
	logFile *os.File
}

func UseBasicLogger() *FileLogger {
	if fileLogger == nil {
		log.Fatal("FileLogger is not initialized")
	}
	return fileLogger
}

func InitFileLogger(appName string, filename string) *FileLogger {
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
		}
	})

	return fileLogger
}

func (l *FileLogger) Shutdown() {
	l.logFile.Close()
}

func (l *FileLogger) Info(message string, ctx LogContext) {
	l.logger.Printf("[INFO] %s; CONTEXT: %v\n", message, ctx)
}

func (l *FileLogger) Error(message string, err error, ctx LogContext) {
	l.logger.Printf("[ERROR]: %s; GO_MESSAGE: %s; CONTEXT: %v\n", message, err.Error(), ctx)
}

func (l *FileLogger) Warning(message string, ctx LogContext) {
	l.logger.Printf("[WARNING] %s; CONTEXT: %v\n", message, ctx)
}
