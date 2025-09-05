package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	location := &args[0]

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
