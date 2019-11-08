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
	RequestHash        string     `json:"request_hash"`
	RequestCached      bool       `json:"request_cached"`
	RequestCacheExpiry int        `json:"request_cache_expiry"`
	EpisodesLastPage   int        `json:"episodes_last_page"`
	Episodes           []*Episode `json:"episodes"`
}

// TODO: maybe take in the page number and handle retries higher in the stack?
func (c *Client) GetEpisodes(animeID int) ([]*Episode, error) {
	page := 1
	episodes := []*Episode{}

	for {
		url := buildUrl(fmt.Sprintf("anime/%v/episodes/%d", animeID, page))

		resp, err := c.Get(url)
		if err != nil {
			return []*Episode{}, err
		}
		defer resp.Body.Close()

		episodesResponse := &EpisodesResponse{}
		if err := json.NewDecoder(resp.Body).Decode(episodesResponse); err != nil {
			return []*Episode{}, err
		}

		if resp.StatusCode != 200 {
			return []*Episode{}, fmt.Errorf("received status code %d when requesting Anime with ID %d", resp.StatusCode, animeID)
		}

		episodes = append(episodes, episodesResponse.Episodes...)

		if page <= episodesResponse.EpisodesLastPage {
			page++
		} else {
			break
		}
	}

	return episodes, nil
}
