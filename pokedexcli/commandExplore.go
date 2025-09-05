package main

import "fmt"

func commandExplore(cfg *config) error {
	location := cfg.parameter

	locationResp, err := cfg.pokeapiClient.PokemonList(location)
	if err != nil {
		return err
	}

	fmt.Println("Founded Pokemon:")
	for _, loc := range locationResp.PokemonEncounters {
		fmt.Println(loc.Pokemon.Name)
	}
	return nil

}
