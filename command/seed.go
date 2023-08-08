package command

import (
	"flag"
	"strings"

	"github.com/rizalarfiyan/sql-seeder/constants"
)

type SeedCommand struct{}

func (*SeedCommand) Help() string {
	helpText := `
Usage: sql-seeder seed [options] ...

  Seeder the database with selected file

Options:

  -config=dbconfig.yml   Configuration file to use.
  -env="development"     Environment.
  name                   The name of the seeder
`
	return strings.TrimSpace(helpText)
}

func (*SeedCommand) Synopsis() string {
	return "Seeder the database with selected file"
}

func (c *SeedCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("new", flag.ContinueOnError)
	cmdFlags.Usage = func() {
		constants.UI.Output(c.Help())
	}

	return 0
}
