package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/rafalb8/renv/internal"
	"github.com/rafalb8/renv/internal/log"
)

type ApplyCmd struct {
	File string `arg:"" optional:"" help:"File to apply"`
}

func (a *ApplyCmd) Run() error {
	if a.File == "" {
		log.Debug("Using default file")
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
	log.Info("Applying file", a.File)
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
	log.Debug("Distro check")

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
	log.Debug("Test check", renv.Test)
	return internal.Exec([]string{"sh", "-c", renv.Test}, true)
}

func includes(renv *internal.Renv) error {
	for _, inc := range renv.Include {
		log.Debug("Include", inc)
		inc = path.Join(path.Dir(renv.Path), inc)
		// Ignore error from included files
		err := (&ApplyCmd{File: inc}).Run()
		if err != nil {
			log.Debug(err)
		}
	}
	return nil
}

func packages(renv *internal.Renv) error {
	if len(renv.Packages) == 0 {
		return nil
	}
	log.Debug("Packages", renv.Packages)
	return (&InstallCmd{Packages: renv.Packages}).Run()
}

func files(renv *internal.Renv) error {
	if len(renv.Files) == 0 {
		return nil
	}
	for from, to := range renv.Files {
		log.Debug("Copying file", path.Base(from), "to", to)
		from = path.Join(path.Dir(renv.Path), from)
		err := internal.CopyFile(from, to)
		if err != nil {
			return err
		}
	}
	return nil
}

func cmd(renv *internal.Renv) error {
	if len(renv.CMD) == 0 {
		return nil
	}
	for _, c := range renv.CMD {
		log.Debug("CMD", c)
		err := internal.Exec([]string{"sh", "-c", c})
		if err != nil {
			return err
		}
	}
	return nil
}
