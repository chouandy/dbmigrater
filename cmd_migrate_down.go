package main

import (
	"flag"
	"fmt"
	"strings"

	dbex "github.com/chouandy/go-sdk/db"
)

// MigrateDownCommand the command struct
type MigrateDownCommand struct {
	Number int
}

// Synopsis the synopsis of command
func (c *MigrateDownCommand) Synopsis() string {
	return "Apply 1 or N down migrations"
}

// Help the help of command
func (c *MigrateDownCommand) Help() string {
	helpText := `
Usage: dbmigrater migrate down
	Apply 1 or N down migrations

Options:
  -n    The number of migrations
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *MigrateDownCommand) Run(args []string) int {
	// Init flag
	f := flag.NewFlagSet("db migrate down", flag.ContinueOnError)
	f.IntVar(&c.Number, "n", 1, "n")
	if err := f.Parse(args); err != nil {
		fmt.Println(err)
		return 1
	}

	// New DB Config
	fmt.Print("New DB Config...")
	if err := dbex.NewConfig(); err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	// Migrate Down
	fmt.Print("Migrate Down...")
	if err := MigrateDown(dbex.GetConfig(), c.Number); err != nil {
		if strings.Contains(err.Error(), "file does not exist") {
			fmt.Println("no migrations")
			return 0
		}
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	return 0
}
