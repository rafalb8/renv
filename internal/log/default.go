package log

import (
	"os"

	"github.com/rafalb8/renv/internal/log/level"
)

var defaultLogger = Simple{
	Level: func() level.Level {
		lvl := os.Getenv("GO_LOG")
		if lvl == "" {
			return level.Info
		}
		switch lvl[0] | ' ' {
		case 'd':
			return level.Debug
		case 'i':
			return level.Info
		case 'w':
			return level.Warning
		case 'e':
			return level.Error
		case 'f':
			return level.Fatal
		case 'h':
			return level.Hidden
		default:
			return level.Info
		}
	}(),
}

func Debug(args ...any) {
	defaultLogger.Debug(args...)
}

func Info(args ...any) {
	defaultLogger.Info(args...)
}

func Warning(args ...any) {
	defaultLogger.Warning(args...)
}

func Error(args ...any) {
	defaultLogger.Error(args...)
}

func Fatal(args ...any) {
	defaultLogger.Fatal(args...)
}
