package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getItems(url string) ([]Item, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var items []Item
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, err
	}

	return items, nil
}
