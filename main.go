package main

import (
	_ "github.com/lib/pq"
	"github.com/jkantner2/Gator/internal/config"
	"github.com/jkantner2/Gator/internal/database"
	"database/sql"
	"log"
	"fmt"
	"os"
)

type state struct{
	db		*database.Queries
	cfg		*config.Config
}

func main() {
	cfg, err := config.ReadJSON()
	if err != nil {
		log.Fatalf("Error reading config: %v\n", err)
	}

	db, err := sql.Open("postgres", cfg.Db_url)
	if err != nil {
		log.Fatalf("Error accessing database: %v\n", err)
	}
	dbQueries := database.New(db)

	programState := &state{}
	programState.db = dbQueries
	programState.cfg = &cfg

	programCommands := commands{
		registeredCommands: make(map[string]func(*state, command) error), 
	}
	programCommands.register("login", handlerLogin)
	programCommands.register("register", handlerRegister)
	programCommands.register("reset", handlerReset)
	programCommands.register("users", handlerUsers)
	programCommands.register("agg", handlerAgg)
	programCommands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	programCommands.register("feeds", handlerFeeds)
	programCommands.register("follow", middlewareLoggedIn(handlerFollow))
	programCommands.register("following", middlewareLoggedIn(handlerFollowing))
	programCommands.register("unfollow", handlerRemoveFollow)
	programCommands.register("browse", handlerBrowse)

	userInput := os.Args
	if len(userInput) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
	}

	userCommand := command{
		 command: userInput[1],
		 arguments: userInput[2:],
	}
	err = programCommands.run(programState, userCommand)
	if err != nil {
		log.Fatal(err)
	}
}


