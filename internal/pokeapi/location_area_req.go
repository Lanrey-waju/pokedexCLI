package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (RespShallowLocations, error) {
	endpoint := "/location-area"
	fullURL := baseUrl + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	if val, ok := c.cache.Get(fullURL); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return RespShallowLocations{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationAreasResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return locationAreasResp, err
	}

	c.cache.Add(fullURL, dat)
	return locationAreasResp, nil

}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseUrl + endpoint

	dat, ok := c.cache.Get(fullURL)
	if ok {
		locationArea := LocationArea{}
		err := json.Unmarshal(dat, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationAreasResp := LocationArea{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return locationAreasResp, err
	}

	c.cache.Add(fullURL, dat)
	return locationAreasResp, nil

}
