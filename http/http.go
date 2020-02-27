// Package http implements common http operations.
package http

import (
	"io/ioutil"
	"net/http"
	"time"
)

const clientIdentifier = "origo-test"

var client *http.Client = &http.Client{
	Timeout: 3 * time.Second,
}

func Fetch(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Client-Identifier", clientIdentifier)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
