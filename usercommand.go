package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Long-meow/RSSaggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("Does not have an argument")
	}
	user, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("User %v does not exist", cmd.arguments[0])
	}
	err = s.configDB.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("The uesr has bean set\n")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("Does not have an argument")
	}

	username := cmd.arguments[0]
	now := time.Now()
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      username,
	})
	if err != nil {
		return err
	}
	err = s.configDB.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("New user has bean created\n")
	fmt.Printf("uuid: %v, createdAt: %v, updatedAt: %v, name: %v\n", user.ID, user.CreatedAt, user.UpdatedAt, user.Name)
	return nil
}

func handlerRemoveUser(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("Does not have an argument")
	}

	_, err := s.db.RemoveUser(context.Background(), cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("User %v does not exist", cmd.arguments[0])
	}

	err = s.configDB.SetUser("")
	if err != nil {
		return err
	}
	fmt.Printf("The uesr has bean removed\n")
	return nil
}

func handlerReset(s *state, cmd command) error {

	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return err
	}

	err = s.configDB.SetUser("")
	if err != nil {
		return err
	}
	fmt.Printf("All users have bean removed\n")
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	current_user := s.configDB.UserName
	for _, user := range users {
		if user == current_user {
			fmt.Printf("* %v (current)\n", user)
		} else {
			fmt.Printf("* %v\n", user)
		}
	}
	return nil
}
