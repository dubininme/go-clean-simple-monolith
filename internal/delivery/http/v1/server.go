package v1

import (
	"github.com/labstack/echo/v4"
)

type Server struct {
	echoEngine *echo.Echo
}

func NewServer(e *echo.Echo) *Server {
	return &Server{echoEngine: e}
}

func (s *Server) Start() error {
	return s.echoEngine.Start(":8080")
}
