package main

import (
	"context"
	"fmt"

	"github.com/shri771/Go/BlogAggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	return func(s *state, cmd command) error {

		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUSerName)
		if err != nil {
			return fmt.Errorf("could not Get user %v:", err)
		}

		return handler(s, cmd, user)

	}
}
