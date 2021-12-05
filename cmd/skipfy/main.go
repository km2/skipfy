package main

import (
	"context"
	"fmt"
	"os"

	"github.com/km2/skipfy/internal/app"
)

func run() error {
	app, err := app.NewApp()
	if err != nil {
		return fmt.Errorf("failed to initialize app: %w", err)
	}

	if err := app.Run(context.Background()); err != nil {
		return fmt.Errorf("failed to run app: %w", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
