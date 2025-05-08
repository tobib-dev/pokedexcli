package main

import (
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Enter the pokemon you want to catch!")
	}

	return nil
}
