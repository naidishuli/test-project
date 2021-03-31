package interfaces

type Logger interface {
	SetField(name, value string)
	WithFields(fields map[string]interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
}
