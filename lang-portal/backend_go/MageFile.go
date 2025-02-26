//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
var Default = Init

// Initialize the database
func Init() error {
	fmt.Println("Initializing database...")
	cmd := exec.Command("sh", "-c", "sqlite3 words.db < migrations/0001_init.sql")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Seed the database
func Seed() error {
	fmt.Println("Seeding database...")
	// Example usage of sh package to run a shell command
	if err := sh.Run("echo", "Seeding database..."); err != nil {
		return err
	}
	// Add seeding logic here
	return nil
}

// Migrate the database
func Migrate() error {
	fmt.Println("Migrating database...")
	// Example usage of mg package to run another Mage target
	mg.Deps(Init)
	// Add migration logic here
	return nil
}
