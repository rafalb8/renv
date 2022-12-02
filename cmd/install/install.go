package install

import (
	"fmt"
	"strings"

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
	utils.ElevatePrivilages(append([]string{"install"}, args...)...)

	// Here we're running as root

	pkgMgr := packageManagers[cache.Get[string]("distro")]
	err := utils.RunCommand(fmt.Sprintf("%s %s", pkgMgr, strings.Join(args, " ")))
	if err != nil {
		log.Error("failed to install some packages")
	}
	return true
}

func (cmd CMD) Help() {
	log.Warning("TODO")
}
