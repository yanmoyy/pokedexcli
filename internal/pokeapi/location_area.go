package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationArea(name string) (RespLocationArea, error) {
	url := baseURL + "/location-area/" + name
	if val, ok := c.cache.Get(url); ok {
		locationResp := RespLocationArea{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespLocationArea{}, fmt.Errorf("json Unmarshal: %w", err)
		}
		return locationResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, fmt.Errorf("http NewRequest: %w", err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, fmt.Errorf("httpClient Do: %w", err)
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, fmt.Errorf("io ReadAll: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return RespLocationArea{}, fmt.Errorf("status code (%d) %s", resp.StatusCode, string(dat))
	}

	locationResp := RespLocationArea{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespLocationArea{}, fmt.Errorf("json Unmarshal: %w", err)
	}
	c.cache.Add(url, dat)
	return locationResp, nil
}
