package main

import (
	_ "github.com/chouandy/go-sdk/dotenv/autoload"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"log"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	// New cli
	c := cli.NewCLI("dbmigrater", version)
	// Set args
	c.Args = os.Args[1:]
	// Set commands
	c.Commands = map[string]cli.CommandFactory{
		"create": func() (cli.Command, error) {
			return &CreateCommand{}, nil
		},
		"drop": func() (cli.Command, error) {
			return &DropCommand{}, nil
		},
		"migrate create": func() (cli.Command, error) {
			return &MigrateCreateCommand{}, nil
		},
		"migrate up": func() (cli.Command, error) {
			return &MigrateUpCommand{}, nil
		},
		"migrate down": func() (cli.Command, error) {
			return &MigrateDownCommand{}, nil
		},
		"migrate drop": func() (cli.Command, error) {
			return &MigrateDropCommand{}, nil
		},
	}
	// Run cli
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
