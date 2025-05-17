package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetLocation -
func (c *Client) GetLocation(locationName string) (RespLocationArea, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationResp := RespLocationArea{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespLocationArea{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	locationResp := RespLocationArea{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespLocationArea{}, err
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}
