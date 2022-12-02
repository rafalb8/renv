package types

import (
	"encoding/json"
	"os"
)

type REnv struct {
	Include  []string          // Include other conf files
	Distro   []string          // Check if distroID on list
	Test     string            // Run cmd and check if exited with 0
	Packages []string          // Install pkgs
	CMD      []string          // Run cmd
	Files    map[string]string // Copy files
}

func LoadREnv(path string) *REnv {
	renv := &REnv{}
	f, err := os.Open(path)
	if err != nil {
		return renv
	}
	json.NewDecoder(f).Decode(renv)
	return renv
}
