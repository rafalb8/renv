package log

import (
	"fmt"
	"os"
	"time"

	"github.com/rafalb8/renv/internal/log/level"
)

type Simple struct {
	Level level.Level
}

func (*Simple) format(lvl level.Level, args ...any) {
	msg := fmt.Sprintln(args...)
	fmt.Print(time.Now().Format("15:04:05.00000"), " ", lvl.Color(), " ", msg)
}

func (l *Simple) Debug(args ...any) {
	if l.Level == level.Debug {
		l.format(level.Debug, args...)
	}
}

func (l *Simple) Info(args ...any) {
	if l.Level <= level.Info {
		l.format(level.Info, args...)
	}
}

func (l *Simple) Warning(args ...any) {
	if l.Level <= level.Warning {
		l.format(level.Warning, args...)
	}
}

func (l *Simple) Error(args ...any) {
	if l.Level <= level.Error {
		l.format(level.Error, args...)
	}
}

func (l *Simple) Fatal(args ...any) {
	if l.Level <= level.Fatal {
		l.format(level.Fatal, args...)
	}
	os.Exit(1)
}
