package log

import (
	"fmt"
	"os"
	"time"

	"github.com/rafalb8/renv/logger/level"
)

type SimpleLogger struct {
}

func (log *SimpleLogger) format(lvl level.Level, in ...any) {
	fmt.Print(time.Now().Format("15:04:05.00000"), " ", lvl.Color(), " ")
	fmt.Println(in...)
}

func (log *SimpleLogger) Debug(in ...any) {
	log.format(level.Debug, in...)
}

func (log *SimpleLogger) Info(in ...any) {
	log.format(level.Info, in...)
}

func (log *SimpleLogger) Warning(in ...any) {
	log.format(level.Warning, in...)
}

func (log *SimpleLogger) Error(in ...any) {
	log.format(level.Error, in...)
}

func (log *SimpleLogger) Fatal(in ...any) {
	log.format(level.Fatal, in...)
	os.Exit(1)
}
