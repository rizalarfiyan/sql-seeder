package command

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/rizalarfiyan/sql-seeder/constants"
	"github.com/rizalarfiyan/sql-seeder/utils"
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

	conf := utils.NewConfig()
	conf.SetConfigFlags(cmdFlags)

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	err := conf.LoadEnv()
	if err != nil {
		constants.UI.Error(err.Error())
		return 1
	}

	name := cmdFlags.Arg(0)
	env, err := conf.GetEnv()
	if err != nil {
		constants.UI.Error(err.Error())
		return 1
	}

	err = NewFileSeeder(name, *env)
	if err != nil {
		constants.UI.Error(err.Error())
		return 1
	}

	return 0
}

func NewFileSeeder(name string, env utils.Environment) error {
	state, err := utils.AlphaNumericUnderscore(name)
	if err != nil {
		return err
	}

	if !state {
		return errors.New(`invalid syntax, must be filled with "a-z, A-Z, 0-9, and _"`)
	}

	files, err := utils.ReadDirectory(env.Dir)
	if err != nil {
		return err
	}

	count := 1
	for _, file := range files {
		if strings.HasSuffix(file.Name(), constants.SQLExtension) {
			count++
		}
	}

	fileName := fmt.Sprintf("%s-%s%s", fmt.Sprintf("%05d", count), strings.TrimSpace(name), constants.SQLExtension)
	pathName := path.Join(env.Dir, fileName)
	file, err := os.Create(pathName)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("Success create new file seed: " + fileName)
	return nil
}
