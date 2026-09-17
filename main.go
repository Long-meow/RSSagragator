package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Long-meow/RSSaggregator/internal/config"
	"github.com/Long-meow/RSSaggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	configDB *config.Config
	db       *database.Queries
}

func main() {
	newConfig, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", newConfig.DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	programState := state{configDB: &newConfig, db: dbQueries}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("removeuser", handlerRemoveUser)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerGetFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollowFeed))
	cmds.register("following", middlewareLoggedIn(handlerGetFeedsFollowForUser))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollowFeed))

	if len(os.Args[:]) < 2 {
		log.Fatalf("Usage: command argument")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	err = cmds.run(&programState, command{name: cmdName, arguments: cmdArgs})
	if err != nil {
		log.Fatalf("can not run command %v", err)
	}

}
