package http

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/arimaulana/point-of-no-return/internal/common/pkg/log"
	"github.com/arimaulana/point-of-no-return/internal/sample/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Server struct {
	app *fiber.App
}

func NewHttpServer(db *sqlx.DB, logger log.Logger) (*Server, error) {
	app := fiber.New()

	// implement middleware
	app.Use(cors.New())
	app.Use(recover.New())

	Routing(app, db, logger)

	server := &Server{
		app: app,
	}

	return server, nil
}

func (s *Server) Run(cfg configs.Config, logger log.Logger) error {
	serverErrors := make(chan error, 1)
	port := cfg.AppHttpPort

	// start the service listening for requests
	go func() {
		logger.Infof("http server listening on: %s", port)
		serverErrors <- s.app.Listen(fmt.Sprintf(":%s", port))
	}()

	// shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don'usertransport collect this error.
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	// Blocking main and waiting for shutdown.
	case sig := <-shutdown:
		logger.Infof("start shutdown: %v", sig)

		// // Give outstanding requests a deadline for completion.
		// ctx, cancel := context.WithTimeout(context.Background(), 10)
		// defer cancel()

		// Asking listener to shutdown and load shed.
		// err := s.app.Shutdown(ctx) // not yet implemented on go fiber
		err := s.app.Shutdown()
		if err != nil {
			logger.Infof("graceful shutdown did not complete in %v : %v", 10, err)
			return err
		}

		switch {
		case sig == syscall.SIGSTOP:
			return errors.New("integrity issue caused shutdown")
		case err != nil:
			return errors.Wrap(err, "could not stop server gracefully")
		}

		// clean server

		logger.Infof("server stopped.")
	}

	return nil
}
