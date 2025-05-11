package v1

import (
	"github.com/dubininme/go-clean-simple-monolith/internal/application/usecase/command"
	"github.com/dubininme/go-clean-simple-monolith/internal/config"
	v1 "github.com/dubininme/go-clean-simple-monolith/internal/delivery/http/v1/handler"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/repository"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/service"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(
	orderRepo repository.OrderRepository,
	discountService *service.DiscountService,
	orderEventPublisher *service.OrderEventPublisher,
	cfg *config.Config,
	productSearchRepo repository.ProductSearchRepository,
	markOrderPaidUsecase *command.MarkOrderPaidUsecase,
) *echo.Echo {
	e := echo.New()
	orderHandler := v1.NewOrderHandler(orderRepo, discountService, orderEventPublisher, markOrderPaidUsecase)
	authHandler := v1.NewAuthHandler(cfg.JWTSecret)
	e.POST("/auth/login", authHandler.Login)

	api := e.Group("/api")
	api.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(cfg.JWTSecret),
	}))
	api.GET("/orders/:id", orderHandler.GetOrder)
	api.POST("/orders/:id/pay", orderHandler.MarkPaid)
	productHandler := v1.NewProductHandler(productSearchRepo)
	api.GET("/products/search", productHandler.Search)
	return e
}
