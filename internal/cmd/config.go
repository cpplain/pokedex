package cmd

import (
	"time"

	"github.com/christopherplain/pokedex/internal/api"
)

type Config struct {
	Commands  Commands
	Cache     *api.Cache
	Endpoints api.Endpoints
	Pokedex   api.Pokedex
	LocAreas  *api.LocAreas
	LocArea   *api.LocArea
}

func NewConfig(internal time.Duration) Config {
	config := Config{
		Commands:  newCommands(),
		Cache:     api.NewCache(internal),
		Endpoints: api.NewEndpoints(),
		Pokedex:   api.Pokedex{},
		LocAreas:  &api.LocAreas{},
		LocArea:   &api.LocArea{},
	}

	config.LocAreas.Next = &config.Endpoints.LocAreas

	return config
}
