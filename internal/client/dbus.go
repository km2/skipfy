package client

import (
	"fmt"

	"github.com/dawidd6/go-spotify-dbus"
	"github.com/godbus/dbus"
	"github.com/km2/skipfy/internal/model"
)

type DBusClient struct {
	conn *dbus.Conn
}

func NewDBusClient() (*DBusClient, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return &DBusClient{
		conn: conn,
	}, nil
}

func (c *DBusClient) CurrentTrack() (*model.Track, error) {
	track, err := spotify.GetMetadata(c.conn)
	if err != nil {
		return nil, fmt.Errorf("failed to get current track: %w", err)
	}

	return &model.Track{
		Artists: track.Artist,
		Title:   track.Title,
	}, nil
}

func (c *DBusClient) Skip() error {
	if err := spotify.SendNext(c.conn); err != nil {
		return fmt.Errorf("failed to skip: %w", err)
	}

	return nil
}
