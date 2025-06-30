package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

		if res.StatusCode != 200 {
			return LocationAreaList{}, fmt.Errorf("Statuscode Error: %d", res.StatusCode)
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

func (c *Client) GetLocationArea(name string) (LocationArea, error) {
	url := baseURL + "/location-area/" + name

	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationArea{}, fmt.Errorf("Failed creating request: %v", err)
		}
		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationArea{}, fmt.Errorf("Network Error: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return LocationArea{}, fmt.Errorf("Statuscode Error: %d", res.StatusCode)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationArea{}, fmt.Errorf("Error reading Data: %v", err)
		}
		c.cache.Add(url, data)
	}
	var area LocationArea
	err := json.Unmarshal(data, &area)
	if err != nil {
		return LocationArea{}, fmt.Errorf("Json Unmarshal: %v", err)
	}
	return area, nil
}
