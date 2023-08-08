package command

import (
	"flag"
	"strings"

	"github.com/rizalarfiyan/sql-seeder/constants"
)

type NewCommand struct{}

func (*NewCommand) Help() string {
	helpText := `
Usage: sql-seeder new [options] name

  Create a new a database seeder.

Options:

  -config=dbconfig.yml   Configuration file to use.
  -env="development"     Environment.
  name                   The name of the seeder
`
	return strings.TrimSpace(helpText)
}

func (*NewCommand) Synopsis() string {
	return "Create a new seeder"
}

func (c *NewCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("new", flag.ContinueOnError)
	cmdFlags.Usage = func() {
		constants.UI.Output(c.Help())
	}

	return 0
}
