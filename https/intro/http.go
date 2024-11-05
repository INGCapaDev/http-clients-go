package main

import (
	"fmt"
	"net/url"
)

func errIfNotHTTPS(URL string) error {
	url, err := url.Parse(URL)
	if err != nil {
		return nil
	}
	if url.Scheme != "https" {
		return fmt.Errorf("error: invalid url protocol, only https is allow")
		// return nil
	}
	return nil
}
