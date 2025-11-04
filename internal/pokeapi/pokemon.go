package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		pokemonResponse := Pokemon{}
		err := json.Unmarshal(val, &pokemonResponse)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResponse, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return Pokemon{}, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResponse := Pokemon{}
	err = json.Unmarshal(data, &pokemonResponse)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	return pokemonResponse, nil
}
