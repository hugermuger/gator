package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	names, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list users: %w", err)
	}

	for _, name := range names {
		if name == s.config.CurrentUserName {
			fmt.Printf("%v (current)\n", name)
		} else {
			fmt.Println(name)
		}
	}

	return nil
}
