package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
	"github.com/rizalarfiyan/sql-seeder/command"
	"github.com/rizalarfiyan/sql-seeder/constants"
	"github.com/rizalarfiyan/sql-seeder/utils"
)

func main() {
	os.Exit(mainExec())
}

func mainExec() int {
	constants.UI = &cli.BasicUi{Writer: os.Stdout}

	cli := &cli.CLI{
		Args: os.Args[1:],
		Commands: map[string]cli.CommandFactory{
			"new": func() (cli.Command, error) {
				return &command.NewCommand{}, nil
			},
			"seed": func() (cli.Command, error) {
				return &command.SeedCommand{}, nil
			},
		},
		HelpFunc: cli.BasicHelpFunc("sql-seeder"),
		Version:  utils.GetVersion(),
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}
