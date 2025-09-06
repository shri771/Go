package main

import (
	"fmt"
	"log"

	"github.com/shri771/Go/BlogAggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)
	err = cfg.SetUser("shri")
	if err != nil {
		log.Fatalf("error setting user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Println(cfg)
	fmt.Printf("Read config again: %+v\n", cfg)

}
