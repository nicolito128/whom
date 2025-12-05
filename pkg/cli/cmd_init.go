package cli

import (
	"flag"
	"fmt"

	"github.com/nicolito128/whom/internal/gens"
)

var initCmd = &InitCommand{
	name: "init",
	fset: flag.NewFlagSet("init", flag.ExitOnError),
}

type InitCommand struct {
	name          string
	fset          *flag.FlagSet
	repoName      *string
	repoNameShort *string
}

func (c *InitCommand) Name() string {
	return c.name
}

func (c *InitCommand) FlagSet() *flag.FlagSet {
	return c.fset
}

func (c *InitCommand) Setup() {
	c.repoName = c.fset.String("name", "", "Set the name for a new repository")
	c.repoNameShort = c.fset.String("n", "", "Alias for -name")
}

func (c *InitCommand) Exec() error {
	repoName := *c.repoName
	shortName := *c.repoNameShort
	if repoName == "" && shortName == "" {
		c.fset.Usage()
		return nil
	}
	if repoName != "" && shortName != "" {
		return fmt.Errorf("cannot use both -name and -n flags simultaneously")
	}

	repoName = selectString(repoName, shortName)
	if err := gens.GenerateBaseProject(repoName); err != nil {
		return err
	}

	return nil
}
