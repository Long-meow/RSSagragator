package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Long-meow/RSSaggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, currentUser database.User) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf(" Required two arguments <name> <url>")
	}
	feedId := uuid.New()
	now := time.Now()
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        feedId,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      cmd.arguments[0],
		Url:       cmd.arguments[1],
		UserID:    currentUser.ID,
	})
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("%+v", feed)
	return nil
}

func handlerGetFeeds(s *state, cmd command) error {

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Printf("Name: %v | Url: %v | Username: %v\n", feed.Name, feed.Url, feed.Username)
	}

	return nil
}

func lookupFeed(s *state, url string) (database.Feed, error) {

	feed, err := s.db.LookupFeed(context.Background(), url)
	if err != nil {
		return database.Feed{}, err
	}
	return feed, nil
}

func handlerFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf(" Required an argument <feed url>")
	}

	feed, err := lookupFeed(s, cmd.arguments[0])
	if err != nil {
		return err
	}
	now := time.Now()
	feedFollowRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	fmt.Printf("User: %v | Feed: %v\n", feedFollowRow.Name, feedFollowRow.Name_2)
	return nil
}

func handlerGetFeedsFollowForUser(s *state, cmd command, user database.User) error {

	feeds, err := s.db.GetFeedsFollow(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("User %v followed:\n", user.Name)
	for _, feed := range feeds {
		fmt.Printf("* %v\n", feed)
	}
	return nil
}

func handlerUnfollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf(" Required an argument <feed url>")
	}

	err := s.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		UserID: user.ID,
		Url:    cmd.arguments[0],
	})

	if err != nil {
		return err
	}
	return nil
}
