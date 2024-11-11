package main

import (
	"fmt"
	"net/http"
)

func httpRequest(url string) (*http.Response, error) {
	// Create a new request to be sent to the random dog API endpoint
	req, err := http.NewRequest(http.MethodGet, RANDOM_DOG_API_ENDPOINT, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request to %s", url)
	}

	// Send a request to the random dog API endpoint
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send outbound request to %s", url)
	}
	return resp, nil
}
