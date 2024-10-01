package main

import (
	"fmt"
	"io"
	"net/http"
)

func getItemData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading the response: %w", err)
	}

	return data, nil
}
