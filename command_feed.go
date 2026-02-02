package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hugermuger/gator/internal/database"
)

func handlerAddfeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}

	feedparams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
		Url:       cmd.arguments[1],
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedparams)
	if err != nil {
		return fmt.Errorf("couldn't add feed: %w", err)
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Println("Feed added successfully:")
	printFeed(feed, user.Name)

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	for x, i := range feeds {
		fmt.Printf("Feed #%v/%v\n", x+1, len(feeds))
		user, err := s.db.GetNameByID(context.Background(), i.UserID)
		if err != nil {
			return err
		}
		printFeed(i, user)
	}

	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) == 0 {
		return errors.New("Website needed!")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedfollow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}
	fmt.Println("Feed followed successfully:")
	fmt.Printf(" * Name:    %v\n", feedfollow.FeedName)
	fmt.Printf(" * User:    %v\n", feedfollow.UserName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("Not following any feeds")
		return nil
	}

	for x, i := range feeds {
		fmt.Printf("Feed#%v/%v\n", x+1, len(feeds))
		fmt.Println(i.FeedName)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("You need a url!")
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}
	params := database.UnfollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	err = s.db.Unfollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Println("Successful unfollowed!")

	return nil
}

func printFeed(feed database.Feed, user string) {
	fmt.Printf(" * Name:    %v\n", feed.Name)
	fmt.Printf(" * Url:     %v\n", feed.Url)
	fmt.Printf(" * User:    %v\n", user)
}
