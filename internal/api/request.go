package api

import (
	"io"
	"net/http"
)

type Endpoints struct {
	LocAreas string
	Pokemon  string
}

func NewEndpoints() Endpoints {
	return Endpoints{
		LocAreas: "https://pokeapi.co/api/v2/location-area/",
		Pokemon:  "https://pokeapi.co/api/v2/pokemon/",
	}
}

func HttpGet(url string, cache *Cache) ([]byte, error) {
	if data, ok := cache.get(url); ok {
		return data, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	cache.add(url, body)

	return body, nil
}
