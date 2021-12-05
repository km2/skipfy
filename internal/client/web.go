package client

import (
	"fmt"

	"github.com/km2/skipfy/internal/model"
	"github.com/zmb3/spotify"
)

type WebClient struct {
	client *spotify.Client
}

func NewWebClient(client *spotify.Client) (*WebClient, error) {
	return &WebClient{client: client}, nil
}

func (c *WebClient) CurrentTrack() (*model.Track, error) {
	track, err := c.client.PlayerCurrentlyPlaying()
	if err != nil {
		return nil, fmt.Errorf("failed to get current track: %w", err)
	}

	artists := make([]string, len(track.Item.Artists))
	for i, artist := range track.Item.Artists {
		artists[i] = artist.Name
	}

	return &model.Track{
		Artists: artists,
		Title:   track.Item.Name,
	}, err
}

func (c *WebClient) Skip() error {
	if err := c.client.Next(); err != nil {
		return fmt.Errorf("failed to skip: %w", err)
	}

	return nil
}
