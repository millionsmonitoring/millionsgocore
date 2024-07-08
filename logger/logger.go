package logger

import (
	"io"
	"log/slog"
	"os"

	"github.com/millionsmonitoring/millionsgocore/env"
)

type LogOptions struct {
	RemoveKeys []string   // remove keys from log
	Level      slog.Level // log level
	Writer     io.Writer  // log writing destination
	AddSource  bool
}

type Option func(o *LogOptions)

func Init(opts ...Option) {

	// assign options given
	logOptions := newDefaultLogOptions()
	for _, opt := range opts {
		opt(&logOptions)
	}

	// init logger, set level, set format
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:       logOptions.Level,
		ReplaceAttr: removeKeys(logOptions.RemoveKeys...),
		AddSource:   logOptions.AddSource,
	})
	contextHandler := NewContextHandler(textHandler)
	slog.SetDefault(slog.New(contextHandler))
}

func WithBlacklistKeys(keys ...string) Option {
	return func(o *LogOptions) {
		o.RemoveKeys = keys
	}
}

func WithWriter(writer io.Writer) Option {
	return func(o *LogOptions) {
		o.Writer = writer
	}
}

func DisableSource() Option {
	return func(o *LogOptions) {
		o.AddSource = false
	}
}

func newDefaultLogOptions() LogOptions {
	level := slog.LevelDebug
	if env.IsProduction() {
		level = slog.LevelInfo
	}
	return LogOptions{
		RemoveKeys: []string{},
		Level:      level,
		Writer:     os.Stdout,
		AddSource:  true,
	}
}

// removeKeys returns a function suitable for HandlerOptions.ReplaceAttr
// that removes all Attrs with the given keys.
func removeKeys(keys ...string) func([]string, slog.Attr) slog.Attr {
	return func(groups []string, a slog.Attr) slog.Attr {
		for _, k := range keys {
			if a.Key == k {
				return slog.Attr{}
			}
		}
		return replaceErrorAttr(groups, a)
	}
}
