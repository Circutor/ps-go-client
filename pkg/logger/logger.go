package logger

// LogAdapter struct Logger Interface.
type LogAdapter struct {
	infoLogger Logger
}

// NewLogAdapter init LogAdapter struct.
func NewLogAdapter(infoLogger Logger) *LogAdapter {
	return &LogAdapter{
		infoLogger: infoLogger,
	}
}

// Info adapter func Info method.
func (log *LogAdapter) Info(args ...interface{}) {
	log.infoLogger.Info(args...)
}
