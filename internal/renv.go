package internal

type Renv struct {
	Distro   []string          // /etc/os-release distro names
	Test     string            // Shell boolean expression
	Include  []string          // Paths to other rEnv files
	Packages []string          // Packages to install
	Files    map[string]string // Files to copy
	CMD      []string          // Shell commands to run

	Path string `json:"-"` // Path to current rEnv file
}
