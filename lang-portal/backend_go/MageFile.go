//go:build mage
// +build mage

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/magefile/mage/mg"
	_ "github.com/mattn/go-sqlite3"
)

// Default target to run when none is specified
var Default = Init

// Initialize the database
func Init() error {
	fmt.Println("Initializing database...")

	cmd := exec.Command("sh", "-c", "sqlite3 database/words.db < migrations/0001_init.sql")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error initializing database: %v\n", err)
		return err
	}
	fmt.Println("Database initialized successfully.")
	return nil
}

// Seed the database
func Seed() error {
	fmt.Println("Seeding database...")

	// Open the database
	dbFile := filepath.Join("database", "words.db")
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		return err
	}
	defer db.Close()

	// Read the seed file
	seedFile := filepath.Join("seeds", "words.json")
	data, err := ioutil.ReadFile(seedFile)
	if err != nil {
		fmt.Printf("Error reading seed file: %v\n", err)
		return err
	}

	// Parse the JSON data
	var words []struct {
		French  string `json:"french"`
		English string `json:"english"`
		Parts   string `json:"parts"`
	}
	if err := json.Unmarshal(data, &words); err != nil {
		fmt.Printf("Error parsing JSON data: %v\n", err)
		return err
	}

	// Insert the data into the database
	for _, word := range words {
		_, err := db.Exec("INSERT INTO words (french, english, parts) VALUES (?, ?, ?)", word.French, word.English, word.Parts)
		if err != nil {
			fmt.Printf("Error inserting data into database: %v\n", err)
			return err
		}
	}

	fmt.Println("Database seeded successfully.")
	return nil
}

// Migrate the database
func Migrate() error {
	fmt.Println("Migrating database...")

	// Ensure the database is initialized
	mg.Deps(Init)

	// Get the list of migration files
	migrationFiles, err := filepath.Glob("migrations/*.sql")
	if err != nil {
		fmt.Printf("Error getting migration files: %v\n", err)
		return err
	}

	// Open the database
	dbFile := filepath.Join("database", "words.db")
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		return err
	}
	defer db.Close()

	// Execute each migration file
	for _, file := range migrationFiles {
		fmt.Printf("Running migration: %s\n", file)
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading migration file: %v\n", err)
			return err
		}

		_, err = db.Exec(string(content))
		if err != nil {
			fmt.Printf("Error executing migration file: %v\n", err)
			return err
		}
	}

	fmt.Println("Database migrated successfully.")
	return nil
}
