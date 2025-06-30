package pokeapi

import (
	"encoding/json"
	"fmt"
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
	// url := baseURL + "/location/?offset=1060&limit=20"
	if providedUrl != nil && len(*providedUrl) != 0 {
		url = *providedUrl
	}
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
	decoder := json.NewDecoder(res.Body)
	var areaList LocationAreaList
	err = decoder.Decode(&areaList)
	if err != nil {
		return LocationAreaList{}, fmt.Errorf("Json Decdode: %v", err)
	}
	return areaList, nil
}
