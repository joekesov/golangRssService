package rss

import "net/http"

type Date string

func Read(url string) (*http.Response, error) {
	return ReadWithClient(url, http.DefaultClient)
}

func ReadWithClient(url string, client *http.Client) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
