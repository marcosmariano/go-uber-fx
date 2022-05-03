package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

type GinLogger struct {
	*Logger
}

type FxLogger struct {
	*Logger
}

var (
	globalLogger *Logger
	zapLogger    *zap.Logger
)

func GetLogger() Logger {
	if globalLogger == nil {
		logger := newLoger()
		globalLogger = &logger
	}
	return *globalLogger
}

func (l Logger) GetGinLogger() GinLogger {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)

	return GinLogger{
		Logger: newSugaredLogger(logger),
	}
}

func (l Logger) GetFxLogger() FxLogger {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)

	return FxLogger{
		Logger: newSugaredLogger(logger),
	}
}

func newLoger() Logger {
	config := zap.NewDevelopmentConfig()
	env := os.Getenv("ENV")
	logOutput := os.Getenv("LOG_OUTPUT")

	if env == "development" {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if env == "production" && logOutput != "" {
		config.OutputPaths = []string{logOutput}
	}

	zapLogger, _ = config.Build()
	logger := newSugaredLogger(zapLogger)

	return *logger
}

func newSugaredLogger(logger *zap.Logger) *Logger {
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}

func (l GinLogger) Write(p []byte) (n int, err error) {
	l.Info(string(p))
	return len(p), nil
}

func (l FxLogger) Printf(str string, args ...interface{}) {
	if args == nil {
		l.Info(str)
		return
	}
	l.Infof(str, args)
}
