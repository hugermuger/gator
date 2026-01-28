package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("Username needed!")
	}

	err := s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("Username set to: %v\n", cmd.arguments[0])

	return nil
}
