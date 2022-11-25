package main

import (
	"os"

	log "github.com/rafalb8/renv/logger"
	"github.com/rafalb8/renv/types"
	"gopkg.in/yaml.v3"
)

func main() {
	f, _ := os.ReadFile("example.yaml")

	var mod types.Module = &types.BaseModule{}

	log.Error(yaml.Unmarshal(f, mod))
	log.Info(log.JSON(mod))
}
