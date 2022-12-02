package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	log "github.com/rafalb8/renv/logger"
)

type FIFO []string

func (fifo FIFO) Peek() string {
	if len(fifo) == 0 {
		return ""
	}
	return fifo[0]
}

func (fifo *FIFO) Pop() string {
	if len(*fifo) == 0 {
		return ""
	}
	x := (*fifo)[0]
	*fifo = (*fifo)[1:]
	return x
}

func RunCommand(cmd string) error {
	command := exec.Command("/bin/sh", "-c", cmd)
	command.Env = os.Environ()
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}

func ElevatePrivilages(args ...string) {
	if os.Geteuid() != 0 {
		log.Info("Elevating privileges...")
		exe, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		RunCommand(fmt.Sprintf("sudo -E %s %s", exe, strings.Join(args, " ")))
		os.Exit(0)
	}
}

func RunSelf(args ...string) {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	RunCommand(fmt.Sprintf("%s %s", exe, strings.Join(args, " ")))
}

func SliceContains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func CopyFile(sourceFile, destinationFile string) {
	input, err := os.ReadFile(sourceFile)
	if err != nil {
		log.Error(err)
		return
	}

	err = os.WriteFile(destinationFile, input, 0644)
	if err != nil {
		log.Error(err)
		return
	}
}
