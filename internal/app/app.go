package app

import (
	"context"
	"fmt"
	"time"

	"github.com/km2/skipfy/internal/client"
	"github.com/km2/skipfy/internal/model"
	"github.com/km2/skipfy/internal/skipper"
)

type Client interface {
	CurrentTrack() (*model.Track, error)
	Skip() error
}

type Skipper interface {
	IsSkip(*model.Track) bool
}

type App struct {
	Client       Client
	Skippers     []Skipper
	TimeDuration time.Duration
}

func NewApp() (*App, error) {
	client, err := client.NewDBusClient()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize DBus client: %w", err)
	}

	skippers := []Skipper{
		&skipper.ContainsSkipper{Substr: "instrumental"},
	}

	timeDuration := time.Second

	return &App{
		Client:       client,
		Skippers:     skippers,
		TimeDuration: timeDuration,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	ticker := time.NewTicker(a.TimeDuration)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-ticker.C:
			track, err := a.Client.CurrentTrack()
			if err != nil {
				return fmt.Errorf("failed to get current track: %w", err)
			}

			for _, skipper := range a.Skippers {
				if skipper.IsSkip(track) {
					if err := a.Client.Skip(); err != nil {
						return fmt.Errorf("failed to skip track: %w", err)
					}

				}
			}
		}
	}
}
