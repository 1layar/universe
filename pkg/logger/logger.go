package logger

import (
	"context"
	"log/slog"

	"go-micro.dev/v4/logger"
)

type MicroLogHandler struct {
	group       string
	attrs       []slog.Attr
	levelStatus map[slog.Level]bool
}

// Enabled implements slog.Handler.
func (t *MicroLogHandler) Enabled(ctx context.Context, lv slog.Level) bool {
	t.levelStatus[lv] = true

	return true
}

// Handle implements slog.Handler.
func (t *MicroLogHandler) Handle(ctx context.Context, record slog.Record) error {
	if !t.levelStatus[record.Level] {
		return nil
	}

	switch record.Level {
	case slog.LevelDebug:
		logger.Debug(record.Message)
	case slog.LevelInfo:
		logger.Info(record.Message)
	case slog.LevelWarn:
		logger.Warn(record.Message)
	case slog.LevelError:
		logger.Error(record.Message)
	}

	return nil
}

// WithAttrs implements slog.Handler.
func (t *MicroLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &MicroLogHandler{
		group:       t.group,
		levelStatus: t.levelStatus,
		attrs:       append(t.attrs, attrs...),
	}
}

// WithGroup implements slog.Handler.
func (t *MicroLogHandler) WithGroup(name string) slog.Handler {
	return &MicroLogHandler{
		group:       name,
		levelStatus: t.levelStatus,
		attrs:       t.attrs,
	}
}

func NewMicroLogHandler() *MicroLogHandler {
	return &MicroLogHandler{}
}
