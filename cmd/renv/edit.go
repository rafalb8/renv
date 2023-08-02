package main

import (
	"os"
	"path"

	"github.com/rafalb8/renv/internal"
)

type EditCmd struct {
	
}

func (e *EditCmd) Run() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	return internal.Exec([]string{"code", path.Join(home, ".renv")})
}