package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (respShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return respShallowLocations{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return respShallowLocations{}, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	locationsResp := respShallowLocations{}
	if err := decoder.Decode(&locationsResp); err != nil {
		return respShallowLocations{}, err
	}

	return locationsResp, nil
}
