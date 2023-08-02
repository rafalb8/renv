package main

import (
	"github.com/alecthomas/kong"
)

var cli struct {
	Version VersionCmd `cmd:"" help:"Show version"`
	Apply   ApplyCmd   `cmd:"" help:"Apply rEnv"`
	Edit    EditCmd    `cmd:"" help:"Open VS Code inside config directory"`
	Install InstallCmd `cmd:"" help:"Install packages"`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
