package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/shri771/Go/BlogAggregator/internal/config"
	"github.com/shri771/Go/BlogAggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// Load database
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmd := commands{
		cmds: make(map[string]func(*state, command) error),
	}

	// Register cmds
	cmd.register("login", handlerLogin)
	cmd.register("register", handlerRegister)
	cmd.register("reset", handlerReset)
	cmd.register("users", handlerList)
	cmd.register("agg", handleragg)
	cmd.register("feeds", handlerListFeeds)
	cmd.register("addfeed", middlewareLoggedIn(handlerAddfeed))
	cmd.register("follow", middlewareLoggedIn(handlerFollow))
	cmd.register("following", middlewareLoggedIn(handlerFollwing))
	cmd.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmd.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

	// Create Newclient
	// rssClient := rssapi.NewClient(5 * time.Second)

}
