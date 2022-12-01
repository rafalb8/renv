package main

import (
	"os"
	"strings"

	"github.com/rafalb8/renv/cmd"
	log "github.com/rafalb8/renv/logger"
	"github.com/rafalb8/renv/utils/cache"
)

type REnv struct {
	Include  []string          // Include other conf files
	Distro   []string          // Check if distroID on list
	Test     string            // Run cmd and check if exited with 0
	Packages []string          // Install pkgs
	CMD      []string          // Run cmd
	Files    map[string]string // Copy files
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
