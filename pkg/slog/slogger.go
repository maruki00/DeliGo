//ref: https://josephwoodward.co.uk/2022/11/slog-structured-logging-proposal

package pkgSlog

import (
	"context"
	"log/slog"

	"github.com/sirupsen/logrus"
)

type LogrusHandler struct {
	logger logrus.Logger
}

func (h *LogrusHandler) Enabled(ctx context.Context, _ slog.Level) bool {
	return true
}

func (h *LogrusHandler) Handle(ctx context.Context, rec slog.Record) error {
	fields := make(map[string]interface{}, rec.NumAttrs())

	entry := h.logger.WithFields(fields)

	switch rec.Level {
	case slog.LevelDebug.Level():
		entry.Debug(rec.Message)
	case slog.LevelInfo.Level():
		entry.Info(rec.Message)
	case slog.LevelWarn.Level():
		entry.Warn(rec.Message)
	case slog.LevelError.Level():
		entry.Error(rec.Message)
	}

	return nil
}

func (h *LogrusHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *LogrusHandler) WithGroup(name string) slog.Handler {
	return h
}
