package logger

// Logger interface API for log.Logger.
type Logger interface {
	Info(...interface{})
}

// Func is a bridge between Logger and zap.SugaredLogger.
type Func func(...interface{})

// Info method to print result logger.
func (f Func) Info(args ...interface{}) { f(args...) }
