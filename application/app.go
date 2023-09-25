package application

import (
	"context"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type Application struct{}

func (a *Application) Start(ctx context.Context) error {
	app := fiber.New()
	port := os.Getenv("PORT")

	ch := make(chan error, 1)
	go func() {
		err := app.Listen(port)
		if err != nil {
			ch <- err
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		app.ShutdownWithTimeout(10 * time.Second)
		log.Info().Msg("Server gracefully stopped")
	}

	return nil
}
