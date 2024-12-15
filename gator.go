package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/kourtzaridisr88/gator/internal/config"
	"github.com/kourtzaridisr88/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.Read()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", config.DBUrl)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	state := &State{
		Config: config,
		DB:     dbQueries,
	}

	cmds := registerCommands()
	runInputCommand(state, cmds)
}

func registerCommands() *Commands {
	commands := &Commands{
		cmds: make(map[string]func(*State, []string) error),
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerListUsers)
	commands.register("agg", handlerAggregatorService)
	commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.register("feeds", handlerListFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFollow))
	commands.register("following", middlewareLoggedIn(handlerFollowing))
	commands.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	return commands
}

func runInputCommand(s *State, cmds *Commands) {
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err := cmds.run(s, cmdName, cmdArgs)
	if err != nil {
		log.Fatal(err)
	}
}
