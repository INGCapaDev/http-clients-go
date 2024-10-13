package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getUsers(url string) ([]User, error) {
	var users []User
	resp, err := http.Get(url)
	if err != nil {
		return users, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, fmt.Errorf("error decoding json from response: %w", err)
	}

	return users, nil
}

// Don't touch below this line

type User struct {
	CharacterName string `json:"characterName"`
	Class         string `json:"class"`
	ID            string `json:"id"`
	Level         int    `json:"level"`
	PvpEnabled    bool   `json:"pvpEnabled"`
	User          struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Age      int    `json:"age"`
	} `json:"user"`
}

func logUsers(characters []User) {
	for _, char := range characters {
		fmt.Printf("Character name: %s, Class: %s, Level: %d, User: %s\n", char.CharacterName, char.Class, char.Level, char.User.Name)
	}
}
