package main

import (
	"database/sql"
	"fmt"

	dbex "github.com/chouandy/go-sdk/db"
)

// Statements
var (
	CreateDatabaseStatement = "CREATE DATABASE `%s` DEFAULT CHARACTER SET = '%s' DEFAULT COLLATE '%s';"
	DropDatabaseStatement   = "DROP DATABASE `%s`;"
)

// CreateDatabase create database
func CreateDatabase(config *dbex.Config) error {
	// New db connection
	conn, err := sql.Open(config.Driver, config.DataSourceWithoutDatabase())
	if err != nil {
		return err
	}
	// defer close db connection
	defer conn.Close()
	// New statement
	statement := fmt.Sprintf(CreateDatabaseStatement,
		config.Database, config.Charset, config.DefaultCollate,
	)
	// Exec statement
	if _, err = conn.Exec(statement); err != nil {
		return err
	}

	return nil
}

// DropDatabase drop database
func DropDatabase(config *dbex.Config) error {
	// New db connection
	conn, err := sql.Open(config.Driver, config.DataSourceWithoutDatabase())
	if err != nil {
		return err
	}
	// defer close db connection
	defer conn.Close()
	// New statement
	statement := fmt.Sprintf(DropDatabaseStatement, config.Database)
	// Exec statement
	if _, err = conn.Exec(statement); err != nil {
		return err
	}

	return nil
}
