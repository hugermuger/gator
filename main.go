package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/hugermuger/gator/internal/config"
	"github.com/hugermuger/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

func main() {
	jsonconfig, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	programState := state{
		config: &jsonconfig,
	}

	db, err := sql.Open("postgres", programState.config.DbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	programState.db = dbQueries

	cmds := commands{
		commandmap: map[string]func(*state, command) error{},
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmd := command{
		name:      os.Args[1],
		arguments: os.Args[2:],
	}

	err = cmds.run(&programState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
