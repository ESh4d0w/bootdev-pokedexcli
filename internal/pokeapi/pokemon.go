package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, fmt.Errorf("Failed creating request: %v", err)
		}
		res, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, fmt.Errorf("Network Error: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return Pokemon{}, fmt.Errorf("Statuscode Error: %d", res.StatusCode)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, fmt.Errorf("Error reading Data: %v", err)
		}
		c.cache.Add(url, data)
	}
	var mon Pokemon
	err := json.Unmarshal(data, &mon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Json Unmarshal: %v", err)
	}
	return mon, nil
}
