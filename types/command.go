package types

type Command interface{}
type SimpleCommand string
type DistroCommand struct {
	Distro []Distro
	Cmd    string
}
