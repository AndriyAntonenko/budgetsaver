package logger

type LogContext = interface{}

type Logger interface {
	Info(message string, ctx LogContext)
	Error(message string, err error, ctx LogContext)
	Warning(message string, ctx LogContext)
}
