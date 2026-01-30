package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (respShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	byteBody, ok := c.pokeCache.Get(*pageURL)
	locationsResp := respShallowLocations{}

	if ok {
		json.Unmarshal(byteBody, &locationsResp)
		return locationsResp, nil
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

	byteBody, err = io.ReadAll(resp.Body)
	c.pokeCache.Add(*pageURL, byteBody)

	if err != nil {
		return respShallowLocations{}, err
	}

	json.Unmarshal(byteBody, &locationsResp)
	return locationsResp, nil
}
