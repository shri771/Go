package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonList(location *string) (Pokemon, error) {
	url := baseURL + "/location-area/" + *location

	// Check if the cache exists
	fmt.Println("Exploring pastoria-city-area...")
	data, ok := c.cache.Get(url)
	if !ok {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}

		c.cache.Add(url, data)
	}

	// Decode the data
	pokemonresp := Pokemon{}
	err := json.Unmarshal(data, &pokemonresp)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemonresp, nil

}
