package main

import (
	"fmt"

	"github.com/rafalb8/renv/internal"
	"github.com/rafalb8/renv/internal/log"
)

type packageManager struct {
	Update  []string
	Install []string
	Check   []string
}

// Map distro to package manager
var managers = map[string]packageManager{
	"arch": {
		Update:  []string{"pacman", "-Sy"},
		Install: []string{"pacman", "-S", "--noconfirm"},
		Check:   []string{"pacman", "-Q"},
	},
	"archarm": {
		Update:  []string{"pacman", "-Sy"},
		Install: []string{"pacman", "-S", "--noconfirm"},
		Check:   []string{"pacman", "-Q"},
	},
	"alpine": {
		Update:  []string{"apk", "update"},
		Install: []string{"apk", "add"},
		Check:   []string{"apk", "info", "-e"},
	},
	"fedora": {
		Update:  []string{"dnf", "update"},
		Install: []string{"dnf", "install", "-y"},
		Check:   []string{"dnf", "list", "installed"},
	},
	"termux": {
		Update:  []string{"apt", "update"},
		Install: []string{"apt", "install", "-y"},
		Check:   []string{"sh", "-c", `dpkg-query -l "$1" | grep -wq "^ii"`, "--"},
	},
	"ubuntu": {
		Update:  []string{"apt", "update"},
		Install: []string{"apt", "install", "-y"},
		Check:   []string{"sh", "-c", `dpkg-query -l "$1" | grep -wq "^ii"`, "--"},
	},
	"debian": {
		Update:  []string{"apt", "update"},
		Install: []string{"apt", "install", "-y"},
		Check:   []string{"sh", "-c", `dpkg-query -l "$1" | grep -wq "^ii"`, "--"},
	},
}

func getMissingPackages(pm packageManager, packages []string) []string {
	missing := make([]string, 0, len(packages))
	for _, pkg := range packages {
		err := internal.Exec(append(pm.Check, pkg), true)
		if err != nil {
			missing = append(missing, pkg)
		}
	}
	return missing
}

type InstallCmd struct {
	Packages []string `arg:"" help:"Packages to install"`
}

func (i *InstallCmd) Run() error {
	// Find package manager
	dist := internal.GetDistro()
	pm, ok := managers[dist]
	if !ok {
		return fmt.Errorf("unknown distro: '%s'", dist)
	}

	// Check for missing packages
	missing := getMissingPackages(pm, i.Packages)
	if len(missing) == 0 {
		log.Warning("all packages are already installed")
		return nil
	}

	// Update
	err := internal.EscalatedExec(pm.Update)
	if err != nil {
		return fmt.Errorf("failed to update: %w", err)
	}

	// Install
	err = internal.EscalatedExec(append(pm.Install, missing...))
	if err != nil {
		return fmt.Errorf("failed to install: %w", err)
	}
	return nil
}
