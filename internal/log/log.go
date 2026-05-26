package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logHistorySize = 50
	envLogLevel    = "LOG_LEVEL"
	envLogFormat   = "LOG_FORMAT"
)

var (
	rootLogger *zap.Logger
	logs       []string
)

// GetLogger returns a named logger.
func GetLogger(name string) *zap.SugaredLogger { _ = "STUB: not implemented"; return nil }

func init() {
	logger, err := initRootLogger()
	if err != nil {
		panic(err)
	}
	rootLogger = logger
}

func initRootLogger() (*zap.Logger, error) { _ = "STUB: not implemented"; return nil, nil }

type logList struct {
	zapcore.LevelEnabler
	enc zapcore.Encoder
}

func (l *logList) clone() *logList { _ = "STUB: not implemented"; return nil }

func (l *logList) With(fields []zapcore.Field) zapcore.Core {
	_ = "STUB: not implemented"
	return *new(zapcore.Core)
}

func (l *logList) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}

func (l *logList) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	_ = "STUB: not implemented"
	return nil
}

func (*logList) Sync() error {
	_ = "STUB: not implemented"

	// Logs get the current logs.
	return nil
}

func Logs() []string {
	_ = "STUB: not implemented"

	// Clear  the current logs.
	return nil
}

func Clear() { _ = "STUB: not implemented"; return }

func addFields(enc zapcore.ObjectEncoder, fields []zapcore.Field) {
	_ = "STUB: not implemented"
	return
}
