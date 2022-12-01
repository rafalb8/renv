package main

import (
	"os"
	"strings"

	"github.com/rafalb8/renv/cmd"
	log "github.com/rafalb8/renv/logger"
	"github.com/rafalb8/renv/utils/cache"
)

type REnv struct {
	Include []string
	Bin     []string
	CMD     []string
	Files   map[string]string
}

func init() {
	cache.Set("distro", func() any {
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
	})
}

func main() {
	cmd.Run()
}
