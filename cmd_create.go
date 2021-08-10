package main

import (
	"fmt"
	"strings"

	dbex "github.com/chouandy/go-sdk/db"
)

// CreateCommand the command struct
type CreateCommand struct{}

// Synopsis the synopsis of command
func (c *CreateCommand) Synopsis() string {
	return "Create database"
}

// Help the help of command
func (c *CreateCommand) Help() string {
	helpText := `
Usage: dbmigrater create
  Create database
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *CreateCommand) Run(args []string) int {
	// New DB Config
	fmt.Print("New DB Config...")
	if err := dbex.NewConfig(); err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	// Create Database
	fmt.Print("Create Database...")
	if err := CreateDatabase(dbex.GetConfig()); err != nil {
		if strings.Contains(err.Error(), "database exists") {
			fmt.Println("database already exists")
			return 0
		}
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	return 0
}
