package log

import (
	"os"
)

type Logger interface {
	Debug(...any)
	Info(...any)
	Warning(...any)
	Error(...any)
	Fatal(...any)
}

func init() {
	var logger Logger = &SimpleLogger{}
	minLevel := os.Getenv("GO_LOG")
	if minLevel == "" {
		minLevel = "info"
	}

	switch minLevel[0] | ' ' {
	case 'd':
		Debug = logger.Debug
		fallthrough
	case 'i':
		Info = logger.Info
		fallthrough
	case 'w':
		Warning = logger.Warning
		fallthrough
	case 'e':
		Error = logger.Error
		fallthrough
	case 'f':
		Fatal = logger.Fatal
	}
}

var (
	Debug   func(in ...any) = func(in ...any) {}
	Info    func(in ...any) = func(in ...any) {}
	Warning func(in ...any) = func(in ...any) {}
	Error   func(in ...any) = func(in ...any) {}
	Fatal   func(in ...any) = func(in ...any) {}
)
