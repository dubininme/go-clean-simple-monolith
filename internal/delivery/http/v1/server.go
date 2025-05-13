package v1

import (
	v1 "github.com/dubininme/go-clean-simple-monolith/internal/delivery/http/v1/handler"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echoEngine *echo.Echo
}

func NewServer(
	orderHandler *v1.OrderHandler,
	productHandler *v1.ProductHandler,
	authHandler *v1.AuthHandler,
) *Server {

	e := NewRouter(orderHandler, productHandler, authHandler)
	return &Server{echoEngine: e}
}

func (s *Server) Start() error {
	return s.echoEngine.Start(":8080")
}
