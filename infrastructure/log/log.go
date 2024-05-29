package log

type Log interface {
	Error(err error)
	Warning(message string)
	Info(message string)
}
