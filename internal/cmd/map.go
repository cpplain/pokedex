package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/christopherplain/pokedex/internal/api"
)

func (c Commands) addMapCmd() {
	c["map"] = command{
		use:   "map",
		short: "Displays the next 20 locations",
		run:   c.runMapCmd,
	}
}

func (c Commands) addMapbCmd() {
	c["mapb"] = command{
		use:   "mapb",
		short: "Displays the previous 20 locations",
		run:   c.runMapbCmd,
	}
}

func (c Commands) runMapCmd(arg string, config Config) {
	if config.LocAreas.Next == nil {
		fmt.Println("There are no next locations.")
		return
	}

	printLocAreas(*config.LocAreas.Next, config)
}

func (c Commands) runMapbCmd(arg string, config Config) {
	if config.LocAreas.Previous == nil {
		fmt.Println("There are no previous locations.")
		return
	}

	printLocAreas(*config.LocAreas.Previous, config)
}

func printLocAreas(url string, config Config) {
	body, err := api.HttpGet(url, config.Cache)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &config.LocAreas)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < len(config.LocAreas.Results); i++ {
		fmt.Println(config.LocAreas.Results[i].Name)
	}
}
