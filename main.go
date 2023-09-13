package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/christopherplain/pokedex/internal/cmd"
)

func main() {
	fmt.Println("Welcome to Pokedex!")

	const interval = 5 * time.Minute
	config := cmd.NewConfig(interval)

	for {
		fmt.Print("Pokedex > ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		words := strings.Fields(input)
		if len(words) < 1 {
			continue
		}
		command := words[0]

		var arg string
		if len(words) > 1 {
			arg = words[1]
		}

		config.Commands.Execute(command, arg, config)
	}
}
