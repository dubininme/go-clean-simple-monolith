package v1

import (
	v1 "github.com/dubininme/go-clean-simple-monolith/internal/delivery/http/v1/handler"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(
	orderHandler *v1.OrderHandler,
	productHandler *v1.ProductHandler,
	authHandler *v1.AuthHandler,
) *echo.Echo {
	e := echo.New()

	e.POST("/auth/login", authHandler.Login)

	api := e.Group("/api")
	api.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(cfg.JWTSecret),
	}))

	// Order routes
	api.GET("/orders/:id", orderHandler.GetOrder)
	api.POST("/orders", orderHandler.CreateOrder)
	api.POST("/orders/:id/pay", orderHandler.MarkOrderPaid)

	// Product routes
	api.GET("/products/search", productHandler.Search)

	return e
}
