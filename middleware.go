package main

import (
	"context"

	"github.com/kourtzaridisr88/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *State, args []string, user database.User) error) func(*State, []string) error {
	return func(s *State, args []string) error {
		user, err := s.DB.GetUserByName(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, args, user)
	}
}
