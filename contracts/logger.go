package contracts

type Logger interface {
	WithFields(fields Fields) Logger
	WithField(key string, value interface{}) Logger
	WithError(err error) Logger
	WithException(exception Exception) Logger
	Info(msg string)
	Warn(msg string)
	Debug(msg string)
	Error(msg string)
	Fatal(msg string)
}
