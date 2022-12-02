package types

import (
	"encoding/json"
	"os"
	"path/filepath"

	log "github.com/rafalb8/renv/logger"
	"github.com/rafalb8/renv/utils/cache"
)

type Config struct {
	LastEnvPath string
}

func (cfg *Config) Save() {
	data, err := json.Marshal(cfg)
	if err != nil {
		log.Error(err)
		return
	}

	path := filepath.Join(cache.Get[string]("homedir"), ".config", "renv", "config.json")
	os.MkdirAll(filepath.Dir(path), 0744)
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		log.Error(err)
		return
	}
}
