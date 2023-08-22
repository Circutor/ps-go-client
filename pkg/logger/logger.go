package logger

// LogAdapter struct Logger Interface.
type LogAdapter struct {
	logger Logger
}

// NewLogAdapter init LogAdapter struct.
func NewLogAdapter(infoLogger Logger) *LogAdapter {
	return &LogAdapter{
		logger: infoLogger,
	}
}

// Info adapter func Info method.
func (log *LogAdapter) Info(args ...interface{}) {
	log.logger.Info(args...)
}
