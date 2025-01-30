package flag

import (
	"github.com/spf13/pflag"
)

// Set holds a pflag.FlagSet as well as an exported Name variable for
// allowing improved help usage information.
type Set struct {
	*pflag.FlagSet
	Name string
}

// NewSet returns a new FlagSet for usage in Config objects.
func NewSet(name string) *Set {
	return &Set{
		FlagSet: pflag.NewFlagSet(name, pflag.ContinueOnError),
		Name:    name,
	}
}

type sensitiveString string

func (s *sensitiveString) String() string {
	return "********"
}

func (s *sensitiveString) Set(value string) error {
	*s = sensitiveString(value)
	return nil
}

func (s *sensitiveString) Type() string {
	return "string"
}

func (Set) SensitiveStringVar(p *string, name, value, usage string) {
	*p = value
	pflag.VarP((*sensitiveString)(p), name, "", usage)
}

func (Set) SensitiveStringVarP(p *string, name, shorthand, value, usage string) {
	*p = value
	pflag.VarP((*sensitiveString)(p), name, shorthand, usage)
}
