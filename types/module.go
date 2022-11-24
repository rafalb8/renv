package types

type Module interface{}
type EnabledModule bool
type BaseModule struct {
	Distro    []Distro
	Variables Variables
	Install   []Package
	Modules   []Module
	Run       []Command
}
