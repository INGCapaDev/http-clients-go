package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getItems(url string) ([]Item, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var items []Item
	decoder := json.NewDecoder(res.Body)

	if err = decoder.Decode(&items); err != nil {
		return nil, fmt.Errorf("error decoding response body %w", err)
	}

	return items, nil
}
