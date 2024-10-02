package main

import (
	"encoding/json"
)

func marshalAll[T any](items []T) ([][]byte, error) {
	var jsons [][]byte

	for _, item := range items {
		data, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}

		jsons = append(jsons, data)
	}

	return jsons, nil

}
