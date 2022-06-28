package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Sync()
}

type syncLogger struct {
	Logger
	SyncImplementation func()
}

func (logger syncLogger) Sync() {
	logger.SyncImplementation()
}

func ConfigureDevelopmentLogger(level string) (Logger, error) {
	// configure level
	zapLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		zap.L().Fatal("failed to parse log level")
	}

	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zapLevel,
	), zap.AddCallerSkip(1))
	zap.ReplaceGlobals(logger)
	return syncLogger{SyncImplementation: func() { _ = logger.Sync() }}, nil
}

func ConfigureProductionLogger(level string) (Logger, error) {
	zapLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		zap.L().Fatal("failed to parse log level")
	}

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zapLevel,
	), zap.AddCallerSkip(1))
	zap.ReplaceGlobals(logger)
	return syncLogger{SyncImplementation: func() { _ = logger.Sync() }}, nil
}

func Debug(msg string) {
	zap.L().Debug(msg)
}

func Debugf(template string, args ...interface{}) {
	zap.L().Sugar().Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	zap.L().Sugar().Debugw(msg, keysAndValues...)
}

func Info(msg string) {
	zap.L().Info(msg)
}

func Infof(template string, args ...interface{}) {
	zap.L().Sugar().Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	zap.L().Sugar().Infow(msg, keysAndValues...)
}

func Warn(msg string) {
	zap.L().Warn(msg)
}

func Warnf(template string, args ...interface{}) {
	zap.L().Sugar().Warnf(template, args...)
}

func Error(msg string) {
	zap.L().Warn(msg)
}

func Errorf(template string, args ...interface{}) {
	zap.L().Sugar().Errorf(template, args...)
}

func Fatal(msg string) {
	zap.L().Fatal(msg)
}

func Fatalf(template string, args ...interface{}) {
	zap.L().Sugar().Fatalf(template, args...)
}
