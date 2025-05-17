package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned_input := CleanInput(input)
		fmt.Printf("Your command was: %s\n", cleaned_input[0])
	}
}
