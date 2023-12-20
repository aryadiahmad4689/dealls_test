package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	db "github.com/aryadiahmad4689/dealls_test/src/app/db"
)

func InitSqliteMaster(ctx context.Context) *sql.DB {

	sqliteClient, err := db.Init(ctx, db.Config{
		Driver: os.Getenv("DB_DRIVER"),
		Source: os.Getenv("DB_PATH"),
	})
	if err != nil {
		fmt.Printf("cannot connect to sqlite: %s\n", err.Error())
		os.Exit(1)
	}
	return sqliteClient.(*sql.DB)
}
func InitSqliteSlave(ctx context.Context) *sql.DB {
	sqliteClient, err := db.Init(ctx, db.Config{
		Driver: os.Getenv("DB_DRIVER"),
		Source: os.Getenv("DB_PATH"),
	})
	if err != nil {
		fmt.Printf("cannot connect to sqlite: %s\n", err.Error())
		os.Exit(1)
	}
	return sqliteClient.(*sql.DB)
}
