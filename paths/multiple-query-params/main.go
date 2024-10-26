package main

import (
	"fmt"
	"strconv"
)

type availabilityStatus = string

const (
	START_QP                     = "?"
	SORT_BY_ESTIMATE_QUERY_PARAM = "sort=estimate"
	LIMIT_QUERY_PARAM            = "limit="
	SEPARATOR_QP                 = "&"

	LOW    availabilityStatus = "Low"
	MEDIUM availabilityStatus = "Medium"
	HIGH   availabilityStatus = "High"
)

func fetchTasks(baseURL, availability string) []Issue {

	availabilityValue, err := getAvailabilityValue(availabilityStatus(availability))
	if err != nil {
		return []Issue{}
	}

	fullURL := baseURL + START_QP + SORT_BY_ESTIMATE_QUERY_PARAM + SEPARATOR_QP + LIMIT_QUERY_PARAM + strconv.Itoa(availabilityValue)
	fmt.Println("URL", fullURL)
	return getIssues(fullURL)
}

func getAvailabilityValue(status availabilityStatus) (int, error) {
	availabilityValues := map[availabilityStatus]int{
		LOW:    1,
		MEDIUM: 3,
		HIGH:   5,
	}

	value, ok := availabilityValues[status]
	if !ok {
		return 0, fmt.Errorf("error getting availability value invalid status")
	}

	return value, nil
}
