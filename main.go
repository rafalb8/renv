package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/rafalb8/renv/cmd"
	log "github.com/rafalb8/renv/logger"
	"github.com/rafalb8/renv/types"
	"github.com/rafalb8/renv/utils/cache"
)

func init() {
	cache.Set("homedir", func() any {
		home, _ := os.UserHomeDir()
		return home
	})
	cache.Set("config", func() any {
		cfg := &types.Config{}
		f, err := os.Open(filepath.Join(cache.Get[string]("homedir"), ".config", "renv", "config.json"))
		if err != nil {
			return cfg
		}
		json.NewDecoder(f).Decode(cfg)
		return cfg
	})

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
