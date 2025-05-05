package main

import (
	"fmt"
	"os"
)

func commandExit(*config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
