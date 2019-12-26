package main

import (
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"

	dbex "github.com/chouandy/go-sdk/db"
)

var migrateDir = "db/migrate"

var migratePath = migrateDir + "/"

var migrateSource = "file://" + migrateDir

// MigrateCreate migrate create
func MigrateCreate(name string) error {
	// Create migrate dir
	if err := os.MkdirAll(migratePath, os.ModePerm); err != nil {
		return err
	}
	// New base
	base := migratePath + time.Now().Format("20060102150405") + "_" + name
	// Generate up and down migrations
	if _, err := os.Create(base + ".up.sql"); err != nil {
		return err
	}
	if _, err := os.Create(base + ".down.sql"); err != nil {
		return err
	}

	return nil
}

// MigrateUp migrate up
func MigrateUp(config *dbex.Config, n int) error {
	// Create migrate dir
	if err := os.MkdirAll(migratePath, os.ModePerm); err != nil {
		return err
	}
	// New migrater
	migrater, err := migrate.New(migrateSource, config.DatabaseURL())
	if err != nil {
		return err
	}
	defer migrater.Close()

	// Check current version is dirty or not
	if v, dirty, err := migrater.Version(); err == nil && dirty {
		// Update the version's dirty to false
		if err := migrater.Force(int(v)); err != nil {
			return err
		}
		// Migrate down version
		if err := migrater.Steps(-1); err != nil {
			return err
		}
	}

	// Migrate up
	if n > 0 {
		if err := migrater.Steps(n); err != nil {
			return err
		}
	} else {
		if err := migrater.Up(); err != nil {
			return err
		}
	}

	return nil
}

// MigrateDown migrate down
func MigrateDown(config *dbex.Config, n int) error {
	// Create migrate dir
	if err := os.MkdirAll(migratePath, os.ModePerm); err != nil {
		return err
	}
	// New migrater
	migrater, err := migrate.New(migrateSource, config.DatabaseURL())
	if err != nil {
		return err
	}
	defer migrater.Close()

	// Check current version is dirty or not
	if v, dirty, err := migrater.Version(); err == nil && dirty {
		// Update the version's dirty to false
		if err := migrater.Force(int(v)); err != nil {
			return err
		}
		// Migrate down version
		if err := migrater.Steps(-1); err != nil {
			return err
		}
	}

	// Migrate down
	if n > 0 {
		if err := migrater.Steps(-n); err != nil {
			return err
		}
	} else {
		if err := migrater.Steps(-1); err != nil {
			return err
		}
	}

	return nil
}

// MigrateDrop migrate drop
func MigrateDrop(config *dbex.Config) error {
	// Create migrate dir
	if err := os.MkdirAll(migratePath, os.ModePerm); err != nil {
		return err
	}
	// New migrater
	migrater, err := migrate.New(migrateSource, config.DatabaseURL())
	if err != nil {
		return err
	}
	defer migrater.Close()

	// Migrate drop
	return migrater.Drop()
}
