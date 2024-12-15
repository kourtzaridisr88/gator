package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/kourtzaridisr88/gator/internal/database"
)

func handlerLogin(s *State, args []string) error {
	if len(args) < 1 {
		return errors.New("login method expect one argument username")
	}
	username := args[0]

	user, err := s.DB.GetUserByName(context.Background(), username)
	if err != nil {
		log.Fatal(err)
	}

	s.Config.SetUser(user.Name)

	fmt.Println("User has been set")

	return nil
}

func handlerRegister(s *State, args []string) error {
	if len(args) < 1 {
		return errors.New("register expect one argument [username]")
	}

	username := args[0]

	params := &database.CreateUserParams{
		ID:        uuid.New(),
		Name:      username,
		CreatedAt: time.Now(),
	}

	user, err := s.DB.CreateUser(context.Background(), *params)
	if err != nil {
		return err
	}

	s.Config.SetUser(user.Name)

	fmt.Println("User has been set")
	fmt.Println(user)

	return nil
}

func handlerReset(s *State, args []string) error {
	return s.DB.TruncateUsers(context.Background())
}

func handlerListUsers(s *State, args []string) error {
	users, err := s.DB.ListUsers(context.Background())
	if err != nil {
		return err
	}

	for _, username := range users {
		val := username

		if val == s.Config.CurrentUserName {
			val += " (current)"
		}

		fmt.Println(val)
	}

	return nil
}

func handlerAggregatorService(s *State, args []string) error {
	if len(args) != 1 {
		return errors.New("invalid number of arguments")
	}

	timeBetweenRequests, err := time.ParseDuration(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %s", args[0])

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func handlerAddFeed(s *State, args []string, user database.User) error {
	if len(args) != 2 {
		return errors.New("addfeed expects name and url as arguments")
	}

	params := &database.CreateFeedParams{
		ID:     uuid.New(),
		Name:   args[0],
		Url:    args[1],
		UserID: user.ID,
	}

	feed, err := s.DB.CreateFeed(context.Background(), *params)
	if err != nil {
		return err
	}

	feedFollowsParams := database.CreateFeedFollowParams{
		ID:     uuid.New(),
		FeedID: feed.ID,
		UserID: user.ID,
	}

	_, err = s.DB.CreateFeedFollow(context.Background(), feedFollowsParams)
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}

func handlerListFeeds(s *State, args []string) error {
	if len(args) > 0 {
		return errors.New("feeds command doesnot accept arguments")
	}

	feeds, err := s.DB.ListFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("User: %v\n", feed.Username)
		fmt.Printf("Name: %v\n", feed.Name)
		fmt.Printf("URL: %v\n", feed.Url)
	}

	return nil
}

func handlerFollow(s *State, args []string, user database.User) error {
	if len(args) != 1 {
		return errors.New("feeds command expects one arguments")
	}

	feed, err := s.DB.GetFeedByUrl(context.Background(), args[0])
	if err != nil {
		return err
	}

	params := database.CreateFeedFollowParams{
		ID:     uuid.New(),
		FeedID: feed.ID,
		UserID: user.ID,
	}

	feedFollow, err := s.DB.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %v\n", feedFollow.FeedName)
	fmt.Printf("User: %v\n", feedFollow.UserName)

	return nil
}

func handlerFollowing(s *State, args []string, user database.User) error {
	if len(args) > 0 {
		return errors.New("following doesnt accept command")
	}

	feeds, err := s.DB.ListFeedsByUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Println("you dont follow any feed")
		return nil
	}

	for _, feedName := range feeds {
		fmt.Printf("User: %s", user.Name)
		fmt.Printf("Feed name: %s", feedName)
	}

	return nil
}

func handlerUnfollow(s *State, args []string, user database.User) error {
	if len(args) != 1 {
		return errors.New("unfollow expect exactly one command")
	}

	feed, err := s.DB.GetFeedByUrl(context.Background(), args[0])
	if err != nil {
		return err
	}

	params := &database.DeleteFeedFollowByUserAndFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.DB.DeleteFeedFollowByUserAndFeed(context.Background(), *params)
	if err != nil {
		return err
	}

	return nil
}
