package application

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/leetcode-golang-classroom/golang-memcache-sample/internal/config"
	"github.com/leetcode-golang-classroom/golang-memcache-sample/internal/memcache"
	"github.com/leetcode-golang-classroom/golang-memcache-sample/internal/types"
	"github.com/leetcode-golang-classroom/golang-memcache-sample/internal/util"
)

type App struct {
	cfg      *config.Config
	appCache types.MemCache
	fiberApp *fiber.App
}

func New(cfg *config.Config) *App {
	// setup fiber
	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: util.DefaultErrorHandler,
	})
	app := &App{
		cfg:      cfg,
		appCache: memcache.Connect(cfg.MemcacheURL),
		fiberApp: fiberApp,
	}
	app.SetupRoutes()
	return app
}

func (app *App) Start(ctx context.Context) error {
	log.Printf("Starting server on %d\n", app.cfg.Port)
	addr := fmt.Sprintf(":%d", app.cfg.Port)
	// errCh channel
	errCh := make(chan error)
	go func() {
		err := app.fiberApp.Listen(addr)
		if err != nil {
			errCh <- fmt.Errorf("failed to start server: %w", err)
		}
		util.CloseChannel(errCh)
	}()
	defer app.Stop()
	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		// cancel from parent
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return app.fiberApp.ShutdownWithContext(timeout)
	}
}

func (app *App) Stop() {
	if err := app.appCache.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to close memory cache %s\n", app.cfg.MemcacheURL)
	}
}
