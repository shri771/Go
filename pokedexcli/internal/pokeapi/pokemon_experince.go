package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Caught struct {
	catched map[string]pokemon
}

type pokemon struct {
}

func (c *Client) Experince(name string) (PokeCatch, error) {

	url := baseURL + "/pokemon/" + name

	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return PokeCatch{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return PokeCatch{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return PokeCatch{}, err
		}

		c.cache.Add(url, data)
	}

	spiceman := PokeCatch{}
	err := json.Unmarshal(data, &spiceman)
	if err != nil {
		return PokeCatch{}, err
	}

	return spiceman, nil

}
