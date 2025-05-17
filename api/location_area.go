package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationAreaBatch struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []locationArea `json:"results"`
}

type locationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func FetchLocationAreasBatch(url string) (*locationAreaBatch, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http get: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code not 200: %v", res.StatusCode)
	}
	var batch locationAreaBatch
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&batch)
	if err != nil {
		return nil, fmt.Errorf("decoder: %w", err)
	}
	return &batch, nil
}
