package database

import (
	"database/sql"
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	DbConnection *sql.DB
)
func DbMigrate(dbParam *sql.DB) {
    migrations := &migrate.FileMigrationSource{
        Dir: "database/sql_migrations",
    }
    n, err := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
    if err != nil {
        panic(err)
    }

	DbConnection = dbParam
    fmt.Println("Applied", n, "migrations")
}