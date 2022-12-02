package edit

import (
	"os"
	"path/filepath"

	log "github.com/rafalb8/renv/logger"
	"github.com/rafalb8/renv/types"
	"github.com/rafalb8/renv/utils"
	"github.com/rafalb8/renv/utils/cache"
)

type CMD struct {
}

func (cmd CMD) Run(args utils.FIFO) bool {
	cfg := cache.Get[*types.Config]("config")

	dir := filepath.Dir(cfg.LastEnvPath)
	if dir == "." {
		dir = filepath.Join(cache.Get[string]("homedir"), ".renv")
		os.MkdirAll(dir, 0744)
	}

	log.Info("Starting vscode in", dir)
	err := utils.RunCommand("code " + dir)
	if err != nil {
		log.Error(err)
	}
	return true
}

func (cmd CMD) Help() {
	log.Info("run vscode inside $HOME/.renv or last env path")
}
