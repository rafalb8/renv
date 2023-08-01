package internal

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func isRoot() bool {
	return os.Geteuid() == 0
}

func Exec(command []string, quiet ...bool) error {
	cmd := exec.Command(command[0], command[1:]...)
	if len(quiet) == 0 || !quiet[0] {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd.Run()
}

func EscalatedExec(command []string, quiet ...bool) error {
	if !isRoot() {
		command = append([]string{"sudo"}, command...)
	}
	return Exec(command, quiet...)
}

func GetDistro() string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		key, value, ok := strings.Cut(line, "=")
		if !ok || key != "ID" {
			continue
		}
		return value
	}
	return ""
}
