package command

import (
	"bufio"
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/rizalarfiyan/sql-seeder/constants"
	"github.com/rizalarfiyan/sql-seeder/utils"
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
	cmdFlags := flag.NewFlagSet("seed", flag.ContinueOnError)
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

	env, err := conf.GetEnv()
	if err != nil {
		constants.UI.Error(err.Error())
		return 1
	}

	db, err := conf.GetConnection(env)
	if err != nil {
		constants.UI.Error(err.Error())
		return 1
	}

	err = NewSeeder(*env, db)
	if err != nil {
		constants.UI.Error(err.Error())
		return 1
	}

	return 0
}

func NewSeeder(env utils.Environment, db *sql.DB) error {
	files, err := utils.ReadDirectory(env.Dir)
	if err != nil {
		return err
	}

	for i, f := range files {
		fmt.Printf("(%d) %s\n", (i + 1), f.Name())
	}

	index, err := InputSelectPrompt("Select file : ")
	if err != nil {
		return err
	}

	if index > len(files) || index < 1 {
		return errors.New("error : bad request file not found")
	}

	file := files[index-1]
	pathName := path.Join(env.Dir, file.Name())
	f, err := os.Open(pathName)
	if err != nil {
		return nil
	}

	bytes := make([]byte, file.Size())
	reader, err := f.Read(bytes)
	if err != nil {
		return nil
	}

	ctx := context.Background()
	query := string(bytes[:reader])
	_, err = db.ExecContext(ctx, query)
	if err != err {
		return err
	}

	return nil
}

func InputSelectPrompt(promptText string) (int, error) {
	readerDb := bufio.NewReader(os.Stdin)
	fmt.Print(promptText)
	text, _ := readerDb.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	i, err := strconv.Atoi(text)
	if err != nil {
		return 0, errors.New("input is not a number")
	}

	return i, nil
}
