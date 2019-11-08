package jikan

import (
	"fmt"
	"net/http"
	"time"
)

const (
	JikanURL = "https://api.jikan.moe/v3"
)

type Client struct {
	*http.Client
}

func New(httpClient *http.Client) *Client {
	c := &Client{
		Client: &http.Client{Timeout: 10 * time.Second},
	}

	if httpClient != nil {
		c.Client = httpClient
	}

	return c
}

func buildUrl(path string) string {
	return fmt.Sprintf("%s/%s", JikanURL, path)
}
