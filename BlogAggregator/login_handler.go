package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("the login handle expects a single argument, the username")
	}

	userName := cmd.args[0]

	err := s.cfg.SetUser(userName)
	if err != nil {
		return errors.New("user registartion failde")
	}
	fmt.Printf("The user: %v has been set", userName)

	return nil
}

type commands struct {
	cmd map[string]func(*state, command) error
}
