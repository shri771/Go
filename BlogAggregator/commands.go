package main

import "github.com/shri771/Go/BlogAggregator/internal/config"

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}
