package install

import (
	"fmt"
	"os/exec"

	log "github.com/rafalb8/renv/logger"
	"github.com/rafalb8/renv/utils"
	"github.com/rafalb8/renv/utils/cache"
)

var packageManagers = map[string]string{
	"arch":    "pacman -Sy --noconfirm",
	"archarm": "pacman -Sy --noconfirm",
	"alpine":  "apk update && apk add",
	"fedora":  "dnf install -y",
	"termux":  "pkg install -y",
	"ubuntu":  "apt update && apt install -y",
	"debian":  "apt update && apt install -y",
}

type CMD struct {
}

func (cmd CMD) Run(args utils.FIFO) bool {
	missing := getMissingPkgs(args)
	if len(missing) == 0 {
		return true
	}

	utils.ElevatePrivilages(append([]string{"install"}, missing...)...)

	// Here we're running as root
	pkgMgr := packageManagers[cache.Get[string]("distro")]

	log.Info("Installing packages:", missing)
	for _, pkg := range missing {
		err := utils.RunCommand(fmt.Sprintf("%s %s", pkgMgr, pkg))
		if err != nil {
			log.Error("failed to install", pkg)
		}
	}
	return true
}

func getMissingPkgs(pkgs []string) []string {
	missing := make([]string, 0, len(pkgs))
	for _, pkg := range pkgs {
		_, err := exec.LookPath(pkg)
		if err != nil {
			missing = append(missing, pkg)
		}
	}
	return missing
}

func (cmd CMD) Help() {
	log.Warning("TODO")
}
