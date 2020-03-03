package main

import (
	"database/sql"

	dbex "github.com/chouandy/go-sdk/db"
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

	// Exec statement
	if _, err = conn.Exec(config.CreateDatabaseStatement()); err != nil {
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

	// Exec statement
	if _, err = conn.Exec(config.DropDatabaseStatement()); err != nil {
		return err
	}

	return nil
}
