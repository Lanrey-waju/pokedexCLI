package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (locationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseUrl + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return locationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return locationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreasResp{}, err
	}

	locationAreasResp := locationAreasResp{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return locationAreasResp, err
	}

	return locationAreasResp, nil

}
