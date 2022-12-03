package apply

import (
	"os"
	"path/filepath"
	"strings"

	log "github.com/rafalb8/renv/logger"
	"github.com/rafalb8/renv/types"
	"github.com/rafalb8/renv/utils"
	"github.com/rafalb8/renv/utils/cache"
)

type CMD struct {
	SaveConfig bool

	env string
}

func (cmd *CMD) Run(args utils.FIFO) bool {
	args = cmd.ParseFlags(args)

	cfg := cache.Get[*types.Config]("config")
	cmd.env = args.Pop()
	if cmd.env == "" {
		// select last applied conf
		cmd.env = cfg.LastEnvPath
		if cmd.env == "" {
			// select default
			cmd.env = cache.Get[string]("defaultEnv")
		}
	}

	if filepath.Ext(cmd.env) != ".json" {
		// set renv.json as env file
		cmd.env = filepath.Join(cmd.env, "renv.json")
	}

	log.Info("Applying file", cmd.env)
	cmd.Execute(types.LoadREnv(cmd.env))

	if cmd.SaveConfig {
		log.Info("Saving config")
		cfg.LastEnvPath, _ = filepath.Abs(cmd.env)
		cfg.Save()
	}
	return true
}

func (cmd CMD) Execute(renv *types.REnv) {
	if renv == nil {
		log.Error("File not found")
		return
	}

	if len(renv.Distro) > 0 {
		log.Info("Checking distro")
		if !utils.SliceContains(renv.Distro, cache.Get[string]("distro")) {
			// skip run
			return
		}
	}

	if renv.Test != "" {
		log.Info("Checking test")
		err := utils.RunCommand(renv.Test)
		if err != nil {
			log.Warning("Test Failed")
			return
		}
		log.Info("Test OK")
	}

	// change dir
	dir := filepath.Dir(cmd.env)
	os.MkdirAll(dir, 0744)
	err := os.Chdir(dir)
	if err != nil {
		log.Error(err)
		return
	}

	for _, include := range renv.Include {
		log.Info("Running include:", include)
		utils.RunSelf("apply", "--child", filepath.Join(filepath.Dir(cmd.env), include))
	}

	// Install packages
	if len(renv.Packages) > 0 {
		utils.RunSelf("install", strings.Join(renv.Packages, " "))
	}

	// Run commands
	for _, command := range renv.CMD {
		utils.RunCommand(command)
	}

	// Copy files
	for src, dst := range renv.Files {
		log.Info("Copying file", src, dst)
		utils.CopyFile(src, dst)
	}
}

func (cmd *CMD) ParseFlags(args utils.FIFO) utils.FIFO {
	for args.Peek() != "" && args.Peek()[0] == '-' {
		switch args.Pop() {
		case "--child":
			cmd.SaveConfig = false
		}
	}
	return args
}

func (cmd CMD) Help() {
	log.Info("possible apply arguments:")
	log.Info("--child: disables config save")
}
