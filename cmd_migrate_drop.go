package main

import (
	"fmt"
	"strings"

	dbex "github.com/chouandy/go-sdk/db"
)

// MigrateDropCommand the command struct
type MigrateDropCommand struct{}

// Synopsis the synopsis of command
func (c *MigrateDropCommand) Synopsis() string {
	return "Drop everything inside database"
}

// Help the help of command
func (c *MigrateDropCommand) Help() string {
	helpText := `
Usage: dbmigrater migrate drop
	Drop everything inside database
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *MigrateDropCommand) Run(args []string) int {
	// New DB Config
	fmt.Print("New DB Config...")
	if err := dbex.NewConfig(); err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	// Migrate Drop
	fmt.Print("Migrate Drop...")
	if err := MigrateDrop(dbex.GetConfig()); err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	return 0
}
