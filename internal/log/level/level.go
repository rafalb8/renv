package level

import "fmt"

type Level uint8

const (
	Debug Level = iota
	Info
	Warning
	Error
	Fatal
	Hidden
)

func (l Level) Color() string {
	color := ""
	switch l {
	case Debug:
		color = "\033[38;5;39m"
	case Info:
		color = "\033[38;5;83m"
	case Warning:
		color = "\033[38;5;220m"
	case Error:
		color = "\033[38;5;196m"
	case Fatal:
		color = "\033[38;5;129m"
	}

	return fmt.Sprintf("%s%s\033[0m", color, l.String())
}

func (l Level) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warning:
		return "WARN"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	default:
		return "-----"
	}
}
