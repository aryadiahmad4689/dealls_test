package db

import (
	"context"
	"errors"

	"github.com/aryadiahmad4689/dealls_test/src/app/db/utils"
)

type Config struct {
	Driver   string
	Source   string
	Host     string
	Password string
	Db       string
}

func Init(ctx context.Context, config Config) (interface{}, error) {
	switch config.Driver {
	case "sqlite3":
		return utils.InitSqlite(config.Driver, config.Source)
	default:
		return nil, errors.New("database driver not found")
	}
}
