package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/rafalb8/renv/internal"
)

type ApplyCmd struct {
	File string `arg:"" optional:"" help:"File to apply"`
}

func (a *ApplyCmd) Run() error {
	if a.File == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		a.File = path.Join(home, ".renv")
	}

	if filepath.Ext(a.File) != ".json" {
		a.File = path.Join(a.File, "renv.json")
	}

	// Read file
	f, err := os.Open(a.File)
	if err != nil {
		return err
	}
	defer f.Close()

	renv := &internal.Renv{
		Path: a.File,
	}
	err = json.NewDecoder(f).Decode(renv)
	if err != nil {
		return err
	}

	executeOrder := []func(*internal.Renv) error{
		distro,
		test,
		includes,
		packages,
		files,
		cmd,
	}

	// Apply
	for _, f := range executeOrder {
		err = f(renv)
		if err != nil {
			return err
		}
	}
	return nil
}

func distro(renv *internal.Renv) error {
	if len(renv.Distro) == 0 {
		return nil
	}
	currentDistro := internal.GetDistro()
	for _, distro := range renv.Distro {
		if distro == currentDistro {
			return nil
		}
	}
	return fmt.Errorf("unknown distro: '%s'", currentDistro)
}

func test(renv *internal.Renv) error {
	if renv.Test == "" {
		return nil
	}
	return internal.Exec([]string{"sh", "-c", renv.Test}, true)
}

func includes(renv *internal.Renv) error {
	for _, f := range renv.Include {
		f := path.Join(path.Dir(renv.Path), f)
		err := (&ApplyCmd{File: f}).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func packages(renv *internal.Renv) error {
	return (&InstallCmd{Packages: renv.Packages}).Run()
}

func files(renv *internal.Renv) error {
	for from, to := range renv.Files {
		err := internal.CopyFile(from, to)
		if err != nil {
			return err
		}
	}
	return nil
}

func cmd(renv *internal.Renv) error {
	for _, f := range renv.CMD {
		err := internal.Exec([]string{"sh", "-c", f})
		if err != nil {
			return err
		}
	}
	return nil
}
