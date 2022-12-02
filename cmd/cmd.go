package cmd

import (
	"os"
	"path/filepath"

	"github.com/rafalb8/renv/cmd/apply"
	"github.com/rafalb8/renv/cmd/edit"
	"github.com/rafalb8/renv/cmd/install"
	log "github.com/rafalb8/renv/logger"
	"github.com/rafalb8/renv/utils"
)

type CMD interface {
	Run(args utils.FIFO) bool
	Help()
}

type REnv struct {
}

func (cmd REnv) Run(args utils.FIFO) bool {
	nextCMD := cmdFromName(args.Pop())
	if !nextCMD.Run(args) {
		cmd.Help()
	}
	return true
}

func (cmd REnv) Help() {
	log.Info("possible renv arguments:")
	log.Info("apply")
	log.Info("edit")
	log.Info("install")
}

type IncorrectCMD string

func (cmd IncorrectCMD) Run(args utils.FIFO) bool {
	cmd.Help()
	return false
}

func (cmd IncorrectCMD) Help() {
	log.Error("incorrect argument:", cmd)
}

func cmdFromName(name string) CMD {
	switch name {
	case "apply":
		return &apply.CMD{SaveConfig: true}
	case "edit":
		return edit.CMD{}
	case "install":
		return install.CMD{}
	case "renv":
		return REnv{}
	}
	return IncorrectCMD(name)
}

func Run() {
	args := utils.FIFO(os.Args)
	cmd := cmdFromName(filepath.Base(args.Pop()))
	cmd.Run(args)
}
