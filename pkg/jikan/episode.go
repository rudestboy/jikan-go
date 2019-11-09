package jikan

import (
	"encoding/json"
	"fmt"
	"time"
)

type Episode struct {
	EpisodeID     int       `json:"episode_id"`
	Title         string    `json:"title"`
	TitleJapanese string    `json:"title_japanese"`
	TitleRomanji  string    `json:"title_romanji"`
	Aired         time.Time `json:"aired"`
	Filler        bool      `json:"filler"`
	Recap         bool      `json:"recap"`
	VideoURL      string    `json:"video_url"`
	ForumURL      string    `json:"forum_url"`
}

type EpisodesResponse struct {
	EpisodesLastPage int        `json:"episodes_last_page"`
	Episodes         []*Episode `json:"episodes"`
}

type GetEpisodeOptions struct {
	Page int
}

func (c *Client) GetEpisodes(animeID int, opts *GetEpisodeOptions) (*EpisodesResponse, error) {
	var page int
	if opts.Page != 0 {
		page = opts.Page
	} else {
		page = 1
	}

	url := buildUrl(fmt.Sprintf("anime/%v/episodes/%d", animeID, page))

	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	episodesResponse := &EpisodesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(episodesResponse); err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("received status code %d when requesting episodes for Anime with ID %d", resp.StatusCode, animeID)
	}

	return episodesResponse, nil
}
