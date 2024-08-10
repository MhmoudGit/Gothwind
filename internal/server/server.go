package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

type Server struct {
	*echo.Echo
	Logger     log.Logger
	ShutdownCh chan os.Signal
	server     *http.Server
}

func New() (*Server, error) {

	e := echo.New()
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt)

	s := &Server{
		Echo:       e,
		Logger:     *log.Default(),
		ShutdownCh: shutdownCh,
	}

	return s, nil
}

// Connect starts the server and listens for shutdown signals
func (s *Server) Connect() {
	s.server = &http.Server{
		Addr:    "localhost:8000",
		Handler: s.RegisterRoutes(),
	}

	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Logger.Fatalf("listen: %s\n", err)
		}
	}()

	s.Logger.Println("Server is running at localhost:8000")
	<-s.ShutdownCh
	s.Shutdown()
}

// Shutdown gracefully shuts down the server and closes the DB connection
func (s *Server) Shutdown() {
	s.Logger.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.Logger.Println("Server forced to shutdown:", err)
	}
	s.Logger.Println("Server exiting")
}
