package server

import (
	"gothwind/cmd/web"
	"gothwind/internal/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	// middlewares
	s.Use(middleware.Logger())
	s.Use(middleware.Recover())
	fileServer := http.FileServer(http.FS(web.Files))
	s.GET("/assets/*", echo.WrapHandler(fileServer))

	s.GET("/", handlers.BaseHandler)

	greet := s.Group("/greet")
	{
		greet.POST("", handlers.GreetPostHandler)
	}

	return s
}
