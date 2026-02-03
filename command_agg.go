package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/hugermuger/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("Please write time in formate 1s, 1m, 1h")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("Please write time in formate 1s, 1m, 1h, Error: %v", err)
	}

	fmt.Printf("Collecting feeds every %v", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	markfeed := database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		ID:        feed.ID,
	}

	err = s.db.MarkFeedFetched(context.Background(), markfeed)
	if err != nil {
		fmt.Println(err)
	}

	onlinefeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Println(err)
	}

	for _, i := range onlinefeed.Channel.Item {
		post := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     i.Title,
			Url:       i.Link,
			Description: sql.NullString{
				String: i.Description,
			},
			PublishedAt: sql.NullString{
				String: i.PubDate,
			},
			FeedID: feed.ID,
		}

		_, err := s.db.CreatePost(context.Background(), post)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(err)
		}
	}
}

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2

	if len(cmd.arguments) != 0 {
		i, err := strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return fmt.Errorf("Argument must be a number")
		}
		limit = i
	}

	feeds, err := s.db.GetFeedFollowForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, f := range feeds {
		params := database.GetPostsForUserParams{
			FeedID: f.FeedID,
			Limit:  int32(limit),
		}

		posts, err := s.db.GetPostsForUser(context.Background(), params)
		if err != nil {
			return err
		}

		fmt.Println()
		fmt.Printf("%v posts from feed %v:\n", limit, f.FeedName)
		fmt.Println()

		for _, p := range posts {
			printPosts(p)
		}

	}

	return nil
}

func printPosts(p database.Post) {
	fmt.Printf("Titel:      %v\n", p.Title)
	fmt.Printf("Url:        %v\n", p.Url)
	fmt.Println(p.Description.String)
	fmt.Println()
}
