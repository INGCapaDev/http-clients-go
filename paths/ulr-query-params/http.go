package main

import (
	"encoding/json"
	"net/http"
)

const (
	SORT_BY_EXP_QUERY_PARAM = "?sort=experience"
)

func getUsers(url string) ([]User, error) {
	fullURL := url + SORT_BY_EXP_QUERY_PARAM
	res, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var users []User
	err = json.NewDecoder(res.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
