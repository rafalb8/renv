package apply

import (
	log "github.com/rafalb8/renv/logger"
	"github.com/rafalb8/renv/utils"
)

type CMD struct {
}

func (cmd CMD) Run(args utils.FIFO) bool {
	cmd.Help()
	return true
}

func (cmd CMD) Help() {
	log.Info("possible apply arguments:")
	log.Warning("TODO")
}
