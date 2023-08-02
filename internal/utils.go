package internal

import (
	"os"
	"os/exec"
	"path"
	"strings"
)

func isRoot() bool {
	return os.Geteuid() == 0
}

func Exec(command []string, quiet ...bool) error {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Env = os.Environ()
	if len(quiet) == 0 || !quiet[0] {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd.Run()
}

func EscalatedExec(command []string, quiet ...bool) error {
	if !isRoot() {
		command = append([]string{"sudo", "-E"}, command...)
	}
	return Exec(command, quiet...)
}

func GetDistro() string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "unknown"
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

func ApplyEnvVars(in string) string {
	// TODO: add support for more variables
	return strings.ReplaceAll(in, "$HOME", os.Getenv("HOME"))
}

func CopyFile(src, dst string) error {
	src = ApplyEnvVars(src)
	dst = ApplyEnvVars(dst)

	err := os.MkdirAll(path.Dir(dst), 0755)
	if err != nil {
		return err
	}
	f, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, f, 0644)
}
