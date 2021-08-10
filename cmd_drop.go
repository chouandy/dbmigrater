package main

import (
	"fmt"
	"strings"

	dbex "github.com/chouandy/go-sdk/db"
)

// DropCommand the command struct
type DropCommand struct{}

// Synopsis the synopsis of command
func (c *DropCommand) Synopsis() string {
	return "Drop database"
}

// Help the help of command
func (c *DropCommand) Help() string {
	helpText := `
Usage: dbmigrater drop
  Drop database
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *DropCommand) Run(args []string) int {
	// New DB Config
	fmt.Print("New DB Config...")
	if err := dbex.NewConfig(); err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	// Drop Database
	fmt.Print("Drop Database...")
	if err := DropDatabase(dbex.GetConfig()); err != nil {
		if strings.Contains(err.Error(), "database doesn't exist") {
			fmt.Println("database doesn't exist")
			return 0
		}
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	return 0
}
