package application

import (
	"context"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/restapi/routes"
)

type Application struct{}

func (a *Application) Start(ctx context.Context) error {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path}\n",
	}))
	port := os.Getenv("PORT")
	routes.LoadAllRoutes(app)

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
