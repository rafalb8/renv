package types

import (
	"encoding/json"
	"os"
	"path/filepath"

	log "github.com/rafalb8/renv/logger"
	"github.com/rafalb8/renv/utils/cache"
)

type Config struct {
	// Path to json file
	LastEnvPath string
}

func (cfg *Config) Save() {
	data, err := json.Marshal(cfg)
	if err != nil {
		log.Error(err)
		return
	}

	path := cache.Get[string]("configPath")
	os.MkdirAll(filepath.Dir(path), 0744)
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		log.Error(err)
		return
	}
}
