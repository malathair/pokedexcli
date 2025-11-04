package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationsResponse, error) {
	url := baseUrl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationsResponse{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return LocationsResponse{}, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return LocationsResponse{}, err
	}

	locationsResponse := LocationsResponse{}
	err = json.Unmarshal(data, &locationsResponse)
	if err != nil {
		return LocationsResponse{}, err
	}

	return locationsResponse, nil
}
