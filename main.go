package main

import (
	"log"
	"os"

	"github.com/hugermuger/gator/internal/config"
)

type state struct {
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

	cmds := commands{
		commandmap: map[string]func(*state, command) error{},
	}

	cmds.register("login", handlerLogin)

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
