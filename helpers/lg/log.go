package lg

import (
	"context"
	"log/slog"
)

// GetLoggerOrDiscard returns the logger if it is not nil, otherwise it returns a discard logger.
func GetLoggerOrDiscard(logger *slog.Logger) *slog.Logger {
	if logger == nil {
		return slog.New(discardHandler{})
	}

	return logger
}

type discardHandler struct {
	disabled bool
}

func (d discardHandler) Enabled(context.Context, slog.Level) bool {
	return !d.disabled
}
func (discardHandler) Handle(context.Context, slog.Record) error {
	return nil
}
func (d discardHandler) WithAttrs([]slog.Attr) slog.Handler {
	return d
}
func (d discardHandler) WithGroup(string) slog.Handler {
	return d
}
