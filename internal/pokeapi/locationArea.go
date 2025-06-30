package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaList struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationAreaList(providedUrl *string) (LocationAreaList, error) {
	url := baseURL + "/location-area"
	if providedUrl != nil && len(*providedUrl) != 0 {
		url = *providedUrl
	}

	data, ok := c.cache.Get(url)

	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaList{}, fmt.Errorf("Failed creating request: %v", err)
		}
		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaList{}, fmt.Errorf("Network Error: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode < 200 && res.StatusCode > 299 {
			return LocationAreaList{}, fmt.Errorf("Statuscode Error: %v", err)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaList{}, fmt.Errorf("Error reading Data: %v", err)
		}
		c.cache.Add(url, data)
	}

	var areaList LocationAreaList
	err := json.Unmarshal(data, &areaList)
	if err != nil {
		return LocationAreaList{}, fmt.Errorf("Json Unmarshal: %v", err)
	}
	return areaList, nil
}
