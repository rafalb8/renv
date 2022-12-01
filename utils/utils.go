package utils

import (
	"os"
	"os/exec"

	log "github.com/rafalb8/renv/logger"
)

type FIFO []string

func (fifo *FIFO) Pop() string {
	if len(*fifo) == 0 {
		return ""
	}
	x := (*fifo)[0]
	*fifo = (*fifo)[1:]
	return x
}

func ElevatePrivilages(args ...string) {
	if os.Geteuid() != 0 {
		log.Info("Elevating privileges...")
		exe, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		cmd := exec.Command("sudo", append([]string{"-E", exe}, args...)...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		os.Exit(0)
	}
}
