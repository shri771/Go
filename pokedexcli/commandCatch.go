package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("enter a argument")
	}

	name := args[0]

	xp, err := cfg.pokeapiClient.Experince(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", name)

	// Genrate a radom number
	maxThrow := 1300
	throw := rand.Intn(maxThrow)

	exp := xp.BaseExperience

	//Check if the catch or miss
	if throw >= exp {
		fmt.Printf("%v was caught!\n", name)
	} else if throw < exp {
		fmt.Printf("%v escaped!\n", name)
	}

	return nil
}
