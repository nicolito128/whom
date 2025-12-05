package cli

import "flag"

type Command interface {
	Name() string
	FlagSet() *flag.FlagSet
	Setup()
	Exec() error
}
