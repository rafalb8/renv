package install

import (
	log "github.com/rafalb8/renv/logger"
	"github.com/rafalb8/renv/utils"
)

type CMD struct {
}

func (cmd CMD) Run(args utils.FIFO) bool {
	utils.ElevatePrivilages(append([]string{"install"}, args...)...)
	cmd.Help()
	return true
}

func (cmd CMD) Help() {
	log.Info("possible install arguments:")
	log.Warning("TODO")
}
