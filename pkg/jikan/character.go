package jikan

import (
	"encoding/json"
	"fmt"
)

type Character struct {
	MalID       int    `json:"mal_id"`
	URL         string `json:"url"`
	ImageURL    string `json:"image_url"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	VoiceActors []struct {
		MalID    int    `json:"mal_id"`
		Name     string `json:"name"`
		URL      string `json:"url"`
		ImageURL string `json:"image_url"`
		Language string `json:"language"`
	} `json:"voice_actors"`
}

type CharactersResponse struct {
	Characters []*Character `json:"characters"`

	// Staff []struct {
	// 	MalID     int      `json:"mal_id"`
	// 	URL       string   `json:"url"`
	// 	Name      string   `json:"name"`
	// 	ImageURL  string   `json:"image_url"`
	// 	Positions []string `json:"positions"`
	// } `json:"staff"`
}

func (c *Client) GetCharacters(animeID int) ([]*Character, error) {
	url := buildUrl(fmt.Sprintf("anime/%v/characters_staff", animeID))

	resp, err := c.Get(url)
	if err != nil {
		return []*Character{}, err
	}
	defer resp.Body.Close()

	charactersResponse := &CharactersResponse{}
	if err := json.NewDecoder(resp.Body).Decode(charactersResponse); err != nil {
		return []*Character{}, err
	}

	return charactersResponse.Characters, nil
}
