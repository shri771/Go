package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.Delusers(context.Background())
	if err != nil {
		return fmt.Errorf("*** Could not reset users table: %w  ****", err)
	}

	err = s.db.DelFeed(context.Background())
	if err != nil {
		return fmt.Errorf("*** Could not reste feeds table: %w ***", err)
	}

	fmt.Printf("Successfully Reseted Users Table")
	return nil
}
