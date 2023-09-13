package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/christopherplain/pokedex/internal/api"
)

func (c Commands) addCatchCmd() {
	c["catch"] = command{
		use:   "catch [pokemon]",
		short: "Tries to catch a Pokemon",
		run:   c.runCatchCmd,
	}
}

func (c Commands) runCatchCmd(pokemonName string, config Config) {
	if _, ok := config.Pokedex[pokemonName]; ok {
		fmt.Println(pokemonName, "already in Pokedex.")
		return
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	url := config.Endpoints.Pokemon + pokemonName + "/"
	body, err := api.HttpGet(url, config.Cache)
	if err != nil {
		fmt.Println(err)
		return
	}

	pokemon := api.Pokemon{}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		fmt.Println(err)
		return
	}

	randNum := rand.Float64()
	catchRate := 5.0 / float64(pokemon.BaseExperience)
	if randNum < catchRate {
		fmt.Println(pokemonName, "caught!")
		config.Pokedex[pokemonName] = pokemon
		return
	}

	fmt.Println(pokemonName, "escaped!")
}
