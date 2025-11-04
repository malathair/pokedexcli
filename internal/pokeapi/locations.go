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

	if val, ok := c.cache.Get(url); ok {
		locationsResponse := LocationsResponse{}
		err := json.Unmarshal(val, &locationsResponse)
		if err != nil {
			return LocationsResponse{}, err
		}
		return locationsResponse, nil
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

	c.cache.Add(url, data)
	return locationsResponse, nil
}

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseUrl + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationResponse := Location{}
		err := json.Unmarshal(val, &locationResponse)
		if err != nil {
			return Location{}, err
		}
		return locationResponse, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return Location{}, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Location{}, err
	}

	locationResponse := Location{}
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)
	return locationResponse, nil
}
