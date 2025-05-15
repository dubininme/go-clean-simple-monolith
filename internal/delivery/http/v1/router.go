package v1

import (
	v1 "github.com/dubininme/go-clean-simple-monolith/internal/delivery/http/v1/handler"

	"github.com/labstack/echo/v4"
)

func NewRouter(
	orderHandler *v1.OrderHandler,
) *echo.Echo {
	e := echo.New()

	api := e.Group("/api/v1")

	// Order routes
	api.GET("/orders/:id", orderHandler.GetOrder)
	api.POST("/orders", orderHandler.CreateOrder)
	api.POST("/orders/:id/pay", orderHandler.MarkOrderPaid)

	return e
}
