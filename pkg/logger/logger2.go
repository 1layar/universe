package logger

import (
	"io"
	"os"
	"testing"

	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

var globalLog *Logger
var zapLogger *zap.Logger

// Logger structure
type Logger struct {
	*zap.SugaredLogger
}

// FxLogger logger for go-fx [subbed from main logger]
type FxLogger struct {
	*Logger
}

// GinLogger logger for gin framework [subbed from main logger]
type GinLogger struct {
	*Logger
}

// GetLogger gets the global instance of the logger
func GetLogger() Logger {
	if globalLog != nil {
		return *globalLog
	}
	globalLog := newLogger()
	return *globalLog
}

// newLogger sets up logger the main logger
func newLogger() *Logger {

	env := os.Getenv("ENVIRONMENT")
	logLevel := os.Getenv("LOG_LEVEL")

	config := zap.NewDevelopmentConfig()

	if env == "local" {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zap.PanicLevel
	}
	config.Level.SetLevel(level)
	zapLogger, _ = config.Build()

	globalLog := zapLogger.Sugar().With(
		zap.String("service", "ppob"),
	)

	return &Logger{
		SugaredLogger: globalLog,
	}

}

func newSugaredLogger(logger *zap.Logger) *Logger {
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}

// GetFxLogger gets logger for go-fx
func (l *Logger) GetFxLogger() fxevent.Logger {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)
	return &FxLogger{Logger: newSugaredLogger(logger)}
}

func (l *FxLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.Logger.Debug("OnStart hook executing: ",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.Logger.Debug("OnStart hook failed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			l.Logger.Debug("OnStart hook executed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.OnStopExecuting:
		l.Logger.Debug("OnStop hook executing: ",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.Logger.Debug("OnStop hook failed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			l.Logger.Debug("OnStop hook executed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.Supplied:
		l.Logger.Debug("supplied: ", zap.String("type", e.TypeName), zap.Error(e.Err))
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.Logger.Debug("provided: ", e.ConstructorName, " => ", rtype)
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			l.Logger.Debug("decorated: ",
				zap.String("decorator", e.DecoratorName),
				zap.String("type", rtype),
			)
		}
	case *fxevent.Invoking:
		l.Logger.Debug("invoking: ", e.FunctionName)
	case *fxevent.Started:
		if e.Err == nil {
			l.Logger.Debug("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err == nil {
			l.Logger.Debug("initialized: custom fxevent.Logger -> ", e.ConstructorName)
		}
	}
}

// GetGinLogger gets logger for gin framework debugging
func (l *Logger) GetGinLogger() io.Writer {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)
	return GinLogger{
		Logger: newSugaredLogger(logger),
	}
}

// Printf prints go-fx logs
func (l FxLogger) Printf(str string, args ...any) {
	if len(args) > 0 {
		l.Debugf(str, args)
	}
	l.Debug(str)
}

// Writer interface implementation for gin-framework
func (l GinLogger) Write(p []byte) (n int, err error) {
	l.Info(string(p))
	return len(p), nil
}

func CreateTestLogger(t *testing.T) Logger {
	zapLogger := zaptest.NewLogger(t)

	sugaredLogger := zapLogger.Sugar()

	return Logger{SugaredLogger: sugaredLogger}
}
