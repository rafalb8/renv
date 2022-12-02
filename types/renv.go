package types

type REnv struct {
	Include  []string          // Include other conf files
	Distro   []string          // Check if distroID on list
	Test     string            // Run cmd and check if exited with 0
	Packages []string          // Install pkgs
	CMD      []string          // Run cmd
	Files    map[string]string // Copy files
}
