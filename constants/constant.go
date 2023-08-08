package constants

import (
	"github.com/go-gorp/gorp/v3"
	"github.com/mitchellh/cli"
)

var UI cli.Ui

var Dialects = map[string]gorp.Dialect{
	"sqlite3":  gorp.SqliteDialect{},
	"postgres": gorp.PostgresDialect{},
	"mysql":    gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"},
}

const SQLExtension = ".sql"
