package cmd

import "fmt"

func (c Commands) addInspectCmd() {
	c["inspect"] = command{
		use:   "inspect [pokemon]",
		short: "Displays details of a Pokemon",
		run:   c.runInspectCmd,
	}
}

func (c Commands) addPokedexCmd() {
	c["pokedex"] = command{
		use:   "pokedex",
		short: "Displays the Pokedex",
		run:   c.runPokedexCmd,
	}
}

func (c Commands) runInspectCmd(pokemonName string, config Config) {
	pokemon, ok := config.Pokedex[pokemonName]
	if !ok {
		fmt.Printf("You have not caught %s.\n", pokemonName)
		return
	}

	fmt.Println("- Name:", pokemon.Name)
	fmt.Println("- Height:", pokemon.Height)
	fmt.Println("- Weight:", pokemon.Weight)
	fmt.Println("- Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("- Types:")
	for _, pType := range pokemon.Types {
		fmt.Println("  -", pType.Type.Name)
	}
}

func (c Commands) runPokedexCmd(arg string, config Config) {
	if len(config.Pokedex) < 1 {
		fmt.Println("Your Pokedex is empty.")
		return
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.Pokedex {
		fmt.Println(" -", pokemon.Name)
	}
}
