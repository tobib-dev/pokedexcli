package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		var command string
		if input.Scan() == true {
			command = input.Text()
		}
		cInput := cleanInput(command)

		fmt.Printf("Your command was: %s\n", cInput[0])
	}
}
