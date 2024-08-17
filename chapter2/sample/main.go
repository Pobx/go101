package main

import (
	"log"
	"os"

	_ "github.com/Pobx/go101/chapter2/sample/matchers"
	"github.com/Pobx/go101/chapter2/sample/search"
)

// func init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
