package initializers

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/millionsmonitoring/millionsgocore/env"
	"github.com/millionsmonitoring/millionsgocore/logger"
)

func InitLogger() *slog.Logger {
	// create a log file for the application
	// if the file already exists, it will append to the file
	// if the file does not exist, it will create a new file
	// the log file will be based on tha name of the environment
	logFile, err := os.OpenFile(fmt.Sprintf("logs/%s.log", env.Env()), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		slog.Error("failed to create log file", slog.String("error", err.Error()))
		panic("failed to create a log file")
	}
	var out io.Writer = logFile

	// adding a logger level for the application
	logLevel := slog.LevelInfo
	if env.IsDevelopment() {
		logLevel = slog.LevelDebug
		out = io.MultiWriter(os.Stdout, logFile)
	}
	h := slog.NewJSONHandler(out, &slog.HandlerOptions{
		ReplaceAttr: logger.ReplaceAttr,
		AddSource:   true,
		Level:       logLevel,
	})
	logger := slog.New(h)
	slog.SetDefault(logger)
	return logger
}
