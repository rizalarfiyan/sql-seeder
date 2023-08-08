package utils

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/rizalarfiyan/sql-seeder/constants"
	"gopkg.in/yaml.v3"
)

type Environment struct {
	Dialect    string `yaml:"dialect"`
	DataSource string `yaml:"datasource"`
	Dir        string `yaml:"dir"`
	TableName  string `yaml:"table"`
}

type Config interface {
	SetConfigFlags(f *flag.FlagSet)
	ReadConfig() (map[string]*Environment, error)
	LoadEnv() error
	GetEnv() (*Environment, error)
	GetConnection(env *Environment) (*sql.DB, error)
}

type config struct {
	file   string
	keyEnv string
	env    *Environment
}

func NewConfig() Config {
	return &config{}
}

func (c *config) SetConfigFlags(f *flag.FlagSet) {
	f.StringVar(&c.file, "config", "dbconfig.yml", "Configuration file to use.")
	f.StringVar(&c.keyEnv, "env", "development", "Environment to use.")
}

func (c *config) ReadConfig() (map[string]*Environment, error) {
	file, err := os.ReadFile(c.file)
	if err != nil {
		return nil, err
	}

	config := make(map[string]*Environment)
	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *config) LoadEnv() error {
	config, err := c.ReadConfig()
	if err != nil {
		return err
	}

	env := config[c.keyEnv]
	if env == nil {
		return fmt.Errorf("no environment: %s", c.keyEnv)
	}

	if env.Dialect == "" {
		return errors.New("no dialect specified")
	}

	if env.TableName == "" {
		return errors.New("no table name specified")
	}

	if env.DataSource == "" {
		return errors.New("no data source specified")
	}

	env.DataSource = os.ExpandEnv(env.DataSource)

	if env.Dir == "" {
		env.Dir = "seeders"
	}

	c.env = env

	if _, err := os.Stat(env.Dir); os.IsNotExist(err) {
		return err
	}

	return nil
}

func (c *config) GetEnv() (*Environment, error) {
	if c.env == nil {
		return nil, errors.New("environment cannot be load")
	}

	return c.env, nil
}

func (c *config) GetConnection(env *Environment) (*sql.DB, error) {
	db, err := sql.Open(env.Dialect, env.DataSource)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to database: %w", err)
	}

	_, exists := constants.Dialects[env.Dialect]
	if !exists {
		return nil, fmt.Errorf("unsupported dialect: %s", env.Dialect)
	}

	return db, nil
}
