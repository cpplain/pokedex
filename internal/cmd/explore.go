package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/christopherplain/pokedex/internal/api"
)

func (c Commands) addExploreCmd() {
	c["explore"] = command{
		use:   "explore [location-area]",
		short: "Displays Pokeman in the given area",
		run:   c.runExploreCmd,
	}
}

func (c Commands) runExploreCmd(locAreaName string, config Config) {
	url := config.Endpoints.LocAreas + locAreaName + "/"
	body, err := api.HttpGet(url, config.Cache)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &config.LocArea)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Exploring %s...\nFound Pokemon:\n", locAreaName)

	for i := 0; i < len(config.LocArea.Encounters); i++ {
		fmt.Printf(" - %s\n", config.LocArea.Encounters[i].Pokemon.Name)
	}
}
